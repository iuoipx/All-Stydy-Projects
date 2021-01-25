#include <stdio.h>
#include <string.h>
#define F(X, Y) X *Y
#define MA(x) x *(x - 1)
#define MIN(x, y) (x) < (y) ? (x) : (y)

int main()
{
    // int a = 3, b = 4;
    // printf("%d\n", F(3 + 5, 3 - 5));

    // int m[][3] = {1, 4, 7, 2, 5, 8, 3, 6, 9};
    // int i, j, k = 2;
    // for (i = 0; i < 3; i++)
    //     printf("%d", m[k][i]);

    // int x = 023;
    // printf("%d\n", --x);

    // int x, y, z;
    // x = 1, y = 10;
    // z = x++, y++, ++y;  //z的值是z = x++ ，如果右边右括号的话，z = ++y
    // printf("%d,%d,%d\n", x, y, z);

    // int x, y, t;
    // x = y = 3;
    // t = ++x || ++y;   //短路， ||运算符左边为真，右边不计算
    // printf("%d", y);

    // int a = 1, b = 5;
    // printf("%d\n", MA(1 + a + b)); //#define 宏就是简单替换，用 1+a+b 去替换x

    // int i = 10, j = 15, k;
    // k = 10 * MIN(i, j);  //宏展开 k=10*(10)<(15)?(10):(15)
    // printf("%d\n", k);

    // int y = 10;
    // while (y--)
    //     ;
    // printf("y=%d\n", y);

    // char str[] = "aeiou", *p = str;
    // printf("%c\n", *p + 4); //先求*p值为a,再+4为c

    // float x = 123.4567;
    // printf("%f\n", (int)(x * 100 + 0.5) / 100.0); //6位有效数字

    // int a = 3, b = 2;
    // b = a < 0 && a++ > 3;  //短路 左边为假值，右边不计算
    // printf("%d,%d\n", a, b);

    // int k = 11;
    // printf("k=%d,k=%o,k=%x\n", k, k, k);  //k=11,k=13,k=b

    // int a = 3, b = 4, c = 5, t = 99;
    // if (b < a && a < c)
    //     t = a;
    // a = c;
    // c = t;
    // if (a < c && b < c)
    //     t = b;
    // b = a;
    // a = t;
    // printf("%d%d%d\n", a, b, c);

    // int a[] = {10, 20, 30, 40, 50, 60}, *p;
    // p = a + 3;
    // printf("%d\n", *p++);

    // char s[] = "ABCD", *p = s;
    // printf("%d\n", *(p + 4));

    // char str[] = "\tab\n\012\\\"";
    // printf("%d\n", strlen(str));

    // char st[80];
    // int i = 0, j = 0;
    // scanf("%s", st);
    // while (st[i])
    // {
    //     if (st[i] >= '0' && st[i] <= '9')
    //         st[j++] = st[i];
    //     i++;
    // }
    // st[j] = '\0';
    // printf("%s\n", st);

    // int a = 3, b = 4, c = 5, d = 2;
    // if (a < b)
    //     if (b < c)
    //         printf("%d", d++ + 1);   //输出d+1=3，此时d值仍然等于2，然后d++，d值为3
    //     else
    //         printf("%d", ++d + 1);
    // printf("%d\n", d);              //输出d的值为3

    // char s[] = "1234567890", *p = s + 3;
    // int i = 5;
    // printf("%d,%s\n", --i, p++);
    // printf("%s\n", p - i);

    // char s[] = "123456789", *p=s, i = 0;
    // while (*p)
    // {
    //     if (i % 2 == 0)
    //         *p = '*';
    //     p++;
    //     i++;
    // }
    // printf("%s\n", s);

    // int a = 0, b = 4, c = 5;
    // switch (a == 0)
    // {
    // case 1:
    //     switch (b < 0)
    //     {
    //     case 1:
    //         printf("@");
    //         break;
    //     case 0:
    //         printf("!");
    //         break;
    //     }
    // case 2:
    //     switch (c == 5)
    //     {
    //     case 0:
    //         printf("*");
    //         break;
    //     case 1:
    //         printf("#");
    //         break;
    //     }
    // default:
    //     printf("&");
    // }

    // int a = 10, *p = &a, *p1;
    // p1 = *&p;
    // printf("%d", *p1);
    // return 0;

    int a[5] = {1, 20, 3, 4, 5};
    int *p = a;
    int num = 6, num1 = 6;
    printf("%d,", *p++);
    printf("%d", *++p);
}