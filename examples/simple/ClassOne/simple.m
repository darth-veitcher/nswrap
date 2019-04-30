#import "simple.h"

@implementation ClassOne

- (instancetype) init
{
	ClassOne *ret;
	ret = [ClassOne alloc];
	ret->i1 = 12;
	ret->p1 = malloc(sizeof(int));
	*ret->p1 = 16;
	ret->a1[0] = 4;
	ret ->a1[1] = 5;
	return ret;
}

- (int) geti1
{
	return i1;
}

- (int *) getp1
{
	return p1;
}

- (int (*)()) getf1
{
	return f;
}

- (int) hi1:(struct stru)in
{
	return in.a;
}

- (int) hi2:(struct stru*)in
{
	return in->a;
}

- (struct stru) nstru1
{
	struct stru ret;
	ret.a = 7;
	ret.b = 8;
	return ret;
}
- (struct stru*) nstru2
{
	struct stru* ret;
	ret = malloc(sizeof(struct stru));
	ret->a = 9;
	ret->b = 10;
	return ret;
}
- (void) hi:(id)in
{
	NSLog(@"hi");
}
- (void) hi3:(id)in
{
	NSLog(@"hi");
}
@end

@implementation ClassTwo
- (instancetype) init
{
	return [super init];
}
@end
