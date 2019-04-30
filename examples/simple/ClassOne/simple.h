#import <Foundation/Foundation.h>

struct stru {int a,b;};

@interface ClassOne : NSObject
{
	int i1;
	int *p1;
	int a1[2];
	int (*f)();
}

- (instancetype) init;
- (int) geti1;
- (int *) getp1;
- (int (*)()) getf1;
- (int) hi1:(struct stru)in;
- (int) hi2:(struct stru*)in;
- (struct stru) nstru1;
- (struct stru*) nstru2;
- (void) hi:(id)in;
- (void) hi3:(id)in;
@end

@interface ClassTwo : ClassOne
{ }
- (instancetype) init;
@end
