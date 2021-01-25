#include <stdio.h>

int main()
{
    float y, res, r = 0.0225;
    printf("please input your money : ");
    scanf("%f", &y);
    res = y * (1 + r);
    printf("your total money : %6.2f", res);
    return 0;
}