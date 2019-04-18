package types

import (
	"fmt"
	"regexp"
)

var (
	whitespace *regexp.Regexp
	wordchars *regexp.Regexp
	reservedwords *regexp.Regexp
)

func init() {
	whitespace = regexp.MustCompile(" *")
	wordchars = regexp.MustCompile(`[_0-9a-zA-Z]`)
	reservedwords = regexp.MustCompile("^(void|char|short|int|long|float|double|signed|unsigned|_Bool|_Complex|const|restrict|volatile|struct|union|enum)$")
}

func dbg(f string, xs ...interface{}) {
	if Debug {
		fmt.Printf(f,xs...)
	}
}

type Parser func(string, *Node) (string, *Node)

// Adders

//Child takes a parser and adds its output node (if non-nil) to the tree.
//FIXME -- broken?
func Child(p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		dbg("Child(%s %p)\n",n.Kind,n)
		s2,n2 := p(s,n)
		if n2 == nil {
			return s,nil
		}
		if n2 != n {
			dbg("Child(%p): AddChild()\n",p)
			n.AddChild(n2)
		}
		return s2,n
	}
}

//ChildOf takes a node and adds results of a parser to it as a child
func ChildOf(ret *Node, p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		dbg("ChildOf(%s %p) %s %p\n",ret.Kind,ret,n.Kind,n)
		s2,n2 := p(s,ret)
		if n2 == nil {
			return s,nil
		}
		if n2 == ret {
			dbg("ChildOf(ret = %p) n2 = %p. WHAT\n",ret,n2)
			ret.Children = n2.Children
		} else {
			dbg("ChildOf(ret = %p) AddChild()\n",ret)
			ret.AddChild(n2)
		}
		return s2,ret
	}
}

//Children takes a parser returns a parser that adds the children of its
//output node to the tree. If multiple parsers are passed in, they are
//passed to Seq(...)
func Children(ps ...Parser) Parser {
	if len(ps) > 1 {
		return Children(Seq(ps...))
	}
	p := ps[0]
	return func(s string, n *Node) (string, *Node) {
		dbg("Children(%s %p)\n",n.Kind,n)
		s2,n2 := p(s,n)
		if n2 == nil {
			return s,nil
		}
		for _,c := range n2.Children {
			dbg("Children(%s %p) AddChild() from %p\n",n.Kind,n,n2)
			if c != n {
				n.AddChild(c)
			}
		}
		return s2,n
	}
}

//ChildrenOf takes a node and adds the children of a parser's output node
//to it as its children.
func ChildrenOf(ret *Node, p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		dbg("ChildrenOf(%s %p) %s %p\n",ret.Kind,ret,n.Kind,n)
		return Children(p)(s,ret)
	}
}

func NodeNamed(k string, p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		s2,n2 := p(s,n)
		if n2 != nil {
			n2.Kind = k
		}
		return s2,n2
	}
}

// Combinators

//Id is the identity parser
func Id(s string, n *Node) (string, *Node) {
	return s,n
}

//Opt optionally runs a Parser, returning the input node if it fails
func Opt(p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		s2,n2 := p(s,n)
		if n2 == nil {
			return s,n
		}
		return s2,n2
	}
}

//OneOf picks the first matching parser and returns its result
func OneOf(ps ...Parser) Parser {
	dbg("OneOf(%p)\n",ps)
	return func(s string, n *Node) (string, *Node) {
		for _,p := range ps {
			s2,n2 := p(s,n)
			if n2 != nil {
				return s2,n2
			}
		}
		return s,nil
	}
}

//Doesn't work? May have side effects that do not get unwound.
func Longest(ps ...Parser) Parser {
	dbg("Longest(%p)\n",ps)
	return func(s string, n *Node) (string, *Node) {
		ss := make([]string,len(ps))
		ns := make([]*Node,len(ps))
		//An arbitrarily large number so I don't have to import "math"
		minrem := 10000
		mini := 0
		for i,p := range ps {
			ss[i],ns[i] = p(s,n)
			if ns[i] != nil && len(ss[i]) < minrem {
				minrem = len(ss[i])
				mini = i
			}
		}
		if minrem < 10000 {
			return ss[mini],ns[mini]
		}
		return s,nil
	}
}

//Seq applies parsers in sequence, adding results as children to the input
//node. Returns nil and the input string unless the entire sequence succeeds
func Seq(ps ...Parser) Parser {
	dbg("Seq(%p)\n",ps)
	return func(s string, n *Node) (string, *Node) {
		ret := NewNode("Seq")
		s2, n2 := s,n
		for _,p := range ps {
			s2, n2 = p(s2,ret)
			if n2 == nil {
				return s,nil
			}
			if n2 != ret {
				dbg("Seq(%p): AddChild()\n",ps)
				ret.AddChild(n2)
			}
		}
		return s2,ret
	}
}
func SeqC(ps ...Parser) Parser {
	return Children(Seq(ps...))
}

//Like Seq but subsequent children are nested inside their earlier siblings.
func Nest(ps ...Parser) Parser {
	dbg("Nest(%p)\n",ps)
	return func(s string, n *Node) (string, *Node) {
		s2,n2 := Seq(ps...)(s,n)
		if n2 == nil {
			return s,nil
		}
		ret := NewNode("Nest")
		n3 := ret
		for _,c := range n2.Children {
			n3.AddChild(c)
			n3 = c
		}
		return s2,ret
	}
}

//ZeroOrMore returns a sequence of zero or more nodes
func ZeroOrMore(p Parser) Parser {
	return func(s string, n *Node) (string, *Node) {
		ret := NewNode("ZeroOrMore")
		dbg("ZeroOrMore(%s %p) ret = %p\n",n.Kind,n,ret)
		var s2 string
		var n2 *Node
		for s2,n2 = p(s,n); n2 != nil; s2,n2 = p(s2,n) {
			dbg("ZeroOrMore(%p): AddChild()\n",p)
			ret.AddChild(n2)
		}
		if len(ret.Children) > 0 {
			return s2,ret
		}
		return s,n
	}
}

func OneOrMore(p Parser) Parser {
	return Seq(p,Children(ZeroOrMore(p)))
}

func Parenthesized(p Parser) Parser {
	return Children(Seq(Lit("("),p,Lit(")")))
}

func Bracketed(p Parser) Parser {
	return Seq(Lit("["),p,Lit("]"))
}

func AngBracketed(p Parser) Parser {
        return Children(Seq(Lit("<"),p,Lit(">")))
}

func CurlyBracketed(p Parser) Parser {
        return Children(Seq(Lit("{"),p,Lit("}")))
}

// Recognizers

func Word(f string) Parser {
	return Lit(f,true)
}

func Lit(f string, ws ...bool) Parser {
	word := false
	if len(ws) > 0 {
		word = ws[0]
	}
	lenf := len(f)
	return func(s string, n *Node) (string, *Node) {
		ret := NewNode("Lit",f)
		dbg("Lit(%p) %s ret = %p\n",n,f,ret)
		if len(s) < lenf {
			return s,nil
		}
		if f == s[:lenf] && !(word && len(s) > lenf && wordchars.Match([]byte{s[lenf]})) {
			adv := lenf
			if loc := whitespace.FindStringIndex(s[lenf:]); loc != nil {
				adv += loc[1]
			}
			return s[adv:],ret
		}
		return s,nil
	}
}

func Regexp(f string) Parser {
	f = "^" + f
	r := regexp.MustCompile(f)
	return func(s string, n *Node) (string, *Node) {
		dbg("Regexp(%p) %s\n",n,f)
		if loc := r.FindStringIndex(s); loc != nil {
			lenf := loc[1]
			adv := lenf
			if loc := whitespace.FindStringIndex(s[lenf:]); loc != nil {
				adv += loc[1]
			}
			ret := NewNode("Regexp",s[:lenf])
			dbg("Regexp(%p): ret = %p (%s)\n",n,ret,s[:lenf])
			return s[adv:],ret
		}
		return s,nil
	}
}

