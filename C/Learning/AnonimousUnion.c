
// C Program to demonstrate working of anonymous union
#include <stdio.h>
struct Scope {
    // Anonymous union
    union {
        char alpha;
        int num;
    };
};
 
int main()
{
    struct Scope x, y;
    x.num = 65;
    y.alpha = 'A';
 
    // Note that members of union are accessed directly
    printf("y.alpha = %c, x.num = %d", y.num, x.num);
 
    return 0;
}