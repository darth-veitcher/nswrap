package types

import (
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

//Parser is a function that takes the string to be parsed plus an input Node
//and returns a new Node and the unparsed remainder string. If the parser fails
//to parse anything in the input, it should return a nil Node.
type Parser func(string, *Node) (string, *Node)

// Adders -- add elements to the Node tree

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

// Combinators -- combine one or more Parsers into a new Parser.

//Opt optionally runs a Parser, returning the input Node (instead of nil)
//if it fails
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

//Seq applies parsers in sequence, adding results as children to the input
//node. Returns nil and the input string unless the entire sequence succeeds
func Seq(ps ...Parser) Parser {
	dbg("Seq(%p)\n",ps)
	p := func(s string, n *Node) (string, *Node) {
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
	return Children(p)
}

//Nest is like Seq but subsequent children are nested inside their earlier
//siblings.
func Nest(ps ...Parser) Parser {
	dbg("Nest(%p)\n",ps)
	p := func(s string, n *Node) (string, *Node) {
		ret := NewNode("Nest")
		s2,n2 := Seq(ps...)(s,ret)
		if n2 == nil {
			return s,nil
		}
		ocs := n2.Children
		ret.Children = []*Node{}
		n3 := ret
		for _,c := range ocs {
			n3.AddChild(c)
			n3 = c
		}
		return s2,ret
	}
	return Children(p)
}

//ZeroOrMore returns a sequence of zero or more nodes
func ZeroOrMore(p Parser) Parser {
	ret := func(s string, n *Node) (string, *Node) {
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
	return Children(ret)
}

//OneOrMore is ZeroOrMore, but fails (returns nil) if the input parser does
//not match any elements.
func OneOrMore(p Parser) Parser {
	return Seq(p,ZeroOrMore(p))
}

//Parenthesized matches the input parser surrounded by literal parenthesis.
func Parenthesized(p Parser) Parser {
	return Children(Seq(Lit("("),p,Lit(")")))
}

//Bracketed matches the input parser surrounded by literal square brackets.
func Bracketed(p Parser) Parser {
	return Seq(Lit("["),p,Lit("]"))
}

//AngBracketed matches the input parser surrounded by literal angled brackets.
func AngBracketed(p Parser) Parser {
        return Children(Seq(Lit("<"),p,Lit(">")))
}

//CurlyBracketed matches the input parser surrounded by literal curly brackets.
func CurlyBracketed(p Parser) Parser {
        return Children(Seq(Lit("{"),p,Lit("}")))
}

// Recognizers -- these functions return parsers that match tokens in the input
// stream. There is no separate tokenizer.

//Word matches an element with a word boundary after its end
func Word(f string) Parser {
	return Lit(f,true)
}

//Lit matches a literal string
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

//Regexp matches a regular expression at the beginning of the input string
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

