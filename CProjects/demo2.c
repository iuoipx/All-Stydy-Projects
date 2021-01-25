#include <stdio.h>
#include <math.h>
#define M 1000

int main()
{
    //4.

    //4.1
    // int a, b, c, max, temp;
    // printf("please input three:");
    // scanf("%d,%d,%d", &a, &b, &c);

    // if (a > b)
    //     if (a > c)
    //         printf("%d", a);
    //     else
    //         printf("%d", c);
    // else if (b > c)
    //     printf("%d", b);
    // else
    //     printf("%d", c);

    //4.2
    // temp = (a > b) ? a : b;
    // max = (temp > c) ? temp : c;
    // printf("max is %d", max);

    //5.
    // int i, k;
    // printf("请输入一个小于%d的正整数:", M);
    // scanf("%d", &i);
    // while (i >= M)
    // {
    //     printf("请重新输入一个小于%d的正整数i:", M);
    //     scanf("%d", &i);
    // }
    // k = sqrt(i);
    // printf("该数的平方根的整数部分是%d", k);

    //6.
    // int x, y;
    // printf("please input x :");
    // scanf("%d", &x);
    // if (x < 1)
    //     y = x;
    // else if (x < 10)
    //     y = 2 * x - 1;
    // else
    //     y = 3 * x - 11;
    // printf("x : %d , y : %d ", x, y);

    //8.
    // float score;
    // printf("please input your gradle: ");
    // scanf("%f", &score);
    // while (score > 100 || score < 0)
    // {
    //     printf("please retry input");
    //     scanf("%f", &score);
    // }
    // switch ((int)(score / 10))
    // {
    // case 10:
    // case 9:
    //     printf("A");
    //     break;
    // case 8:
    //     printf("B");
    //     break;
    // case 7:
    //     printf("C");
    //     break;
    // case 6:
    //     printf("D");
    //     break;
    // default:
    //     printf("E");
    // }

    //9.
    int ch, place, indiv, ten, hundred, thousand, ten_thousand;
    printf("please input :");
    scanf("%d", &ch);
    while (ch > 99999)
    {
        printf("please retry input :");
        scanf("%d", &ch);
    }
    if (ch > 9999)
        place = 5;
    else if (ch > 999)
        place = 4;
    else if (ch > 99)
        place = 3;
    else if (ch > 9)
        place = 2;
    else
        place = 1;

    printf("place : %d", place);
    printf("ch is :");
    ten_thousand = ch / 10000;
    thousand = (ch - ten_thousand * 10000) / 1000;
    hundred = (ch - ten_thousand * 10000 - thousand * 1000) / 100;
    ten = (ch - ten_thousand * 10000 - thousand * 1000 - hundred * 100) / 10;
    indiv = ch - ten_thousand * 10000 - thousand * 1000 - hundred * 100 - ten * 10;
    switch (place)
    {
    case 5:
        printf("%d%d%d%d%d\n", ten_thousand, thousand, hundred, ten, indiv);
        printf("need : %d%d%d%d%d", indiv, ten, hundred, thousand, ten_thousand);
        break;
    case 4:
        printf("%d%d%d%d%d\n", thousand, hundred, ten, indiv);
        printf("need : %d%d%d%d%d", indiv, ten, hundred, thousand);
        break;
    case 3:
        printf("%d%d%d%d%d\n", hundred, ten_thousand, indiv);
        printf("need : %d%d%d%d%d", indiv, ten, hundred);
        break;
    case 2:
        printf("%d%d%d%d%d\n", ten, indiv);
        printf("need : %d%d%d%d%d", indiv, ten);
        break;
    case 1:
        printf("%d%d%d%d%d\n", indiv);
        printf("need : %d%d%d%d%d", indiv);
        break;
    default:
        break;
    }

    return 0;
}