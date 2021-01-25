#include <stdio.h>

int main()
{
    int max4(int a, int b, int c, int d);
    int a, b, c, d, e;
    printf("please input 4 :");
    scanf("%d%d%d%d", &a, &b, &c, &d);
    e = max4(a, b, c, d);
    printf("max is %d", e);
    return 0;
}

int max4(int a, int b, int c, int d)
{
    int max2(int x, int y);

    return max2(max2(max2(a, b), c), d);
}

int max2(int x, int y)
{
    return (x > y ? x : y);
}