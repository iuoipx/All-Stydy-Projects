#include <stdio.h>

int main()
{
    int a, b, c, max;
    printf("please input a,b,c : ");
    scanf("%d%d%d", &a, &b, &c);
    max = (a > b) ? a : b;
    max = (max > c) ? max : c;
    printf("max number : %d", max);
    return 0;
}