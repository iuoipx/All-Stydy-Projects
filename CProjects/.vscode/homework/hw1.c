#include <stdio.h>

int main()
{
    int th, h, t, a, num, newNum;
    printf("please enter a number: ");
    scanf("%d", &num);
    th = num / 1000;
    h = num / 100 % 10;
    t = num / 10 % 10;
    a = num % 10;
    newNum = a * 1000 + t * 100 + h * 10 + th;
    printf("result: %d", newNum);
}