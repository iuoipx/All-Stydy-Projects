#include <stdio.h>

int main()
{

    // int k;
    // float s;
    // for (k = 0, s = 0; k < 7; k++)
    // {
    //     float t = k / 2;
    //     printf("%f \t", t);
    //     s += k / 2;
    // }
    // printf("%d %f", k, s);

    // int a = 7;
    // float x = 2.5, y = 4.7;
    // printf("%f",x + a % 3 * (int)(x + y) % 2 / 4);

    // int num, sum;
    // num = sum = 7;
    // sum = num++, sum++, ++num;
    // printf("%d",sum);

    // int k = 10;
    // while (k = 1)
    //     printf("czx");

    // int n = 0;
    // while (n++ <= 2)
    //     ;
    // printf("%d", n);

    // int num = 0;
    // while (num++ <= 2)
    //     ;

    // printf("%d\n", num);

    // int num = 7, sum = 7;
    // sum = num++, sum++, ++num;
    // printf("%d", (2+2,3+3,4+14,10!=9));

    // int  x,y,z ;
    //  scanf("%d%d%d",&x,&y,&z );
    // printf("x+y+z=%d\n" ,x+y+z);

    //printf("%f",18);

    // int a = 7;
    // float x = 2.5, y = 4.7;
    // printf("%f",x + a % 3 * (int)(x + y) % 2 / 4);

    //printf("%d",sizeof(int));

    // char c = '1' + '1';
    // printf("%d", c);

    // int num, sum;
    // num = sum = 7;
    // sum = num++, sum++, ++num;
    // printf("%d", sum);

    // char c;
    // c=97;
    // printf("%c",c);

    // int m = 3, n = 4, x;
    // x = -m++;
    // printf("%d", x);
    // x = x + 8 / ++n;
    // printf("%d\n", x);

    // int func(char ch, int n);
    // int k = func(97, 22);
    // printf("%d", k);

    // int fun(int a, int b);
    // int x = 3, y = 8, z = 6, r;
    // r = fun(fun(x, y), 2 * z);
    // printf("%d\n", r);

    // char a[]="string";
    // char a1[]={99,97,92,3,4,5};
    // printf("%s",a1);

    // int a = 13, b = 14;
    // int c = a % b;
    // printf("%d", c);

    // char s[10];
    // s="abcd";   //s是常量
    // printf("%s\n",s);
    // return 0;

    // int i, c;
    // int num[][9] = {{1, 2, 3, 4}, {2, 7, 2, 8}};
    // for (i = 0; i < 4; i++)
    // {
    //     c = num[0][i] + num[1][i] - 2 * 0;
    //     printf("%d ", c);
    // }

    // int a1, a2;
    // char c1, c2;
    // scanf("%d%c%d%c", &a1, &c1, &a2, &c2);
    // printf("%d%c%d%c", a1, c1, a2, c2);

    // int a;
    // a++;
    // printf("%d", a);

    //printf("%d", 3 % -2);

    // double x = 1.42, y = 5.2;
    // int i = 8;
    // y = (float)i;
    // printf("%f", y);

    // int a;
    // scanf("%c", &a);
    // printf("%d", a);

    // int w = 3, z = 7, x = 10;
    // printf("%d\n", x > 10 ? x + 100 : x - 10);
    // printf("%d\n", w++ || z++);
    // printf("%d\n", !w > z);
    // printf("%d\n", w && z);

    // int x = 0, s = 0;
    // while (!x != 0)
    //     s += ++x;
    // printf("%d", s);

    // int m = 5;
    // printf("%d", ++m);

    // int i, a[10];
    // for (i = 0; i < 10; i++)
    //     a[i] = i;
    // printf("%d%d%d", a[2], a[5], a[8]);

    // int a, b, temp;
    // printf("请输入两个整数: ");
    // scanf("%d", &a);
    // scanf("%d", &b);
    // if (a > b)
    //     printf("%d,%d", a, b);
    // else
    //     printf("%d,%d", b, a);

    // int res = 1;
    // for (int i = 2; i <= 6; i++)
    //     res *= i;
    // printf("%d", res);

    // int a[8];
    // for (int i = 0; i < 8; i++)
    // {
    //     scanf("%d", &a[i]);
    // }
    // printf("%d", a[7]);

    // char ch;
    // scanf("%c", &ch);
    // if (ch >= 'A' && ch <= 'Z')
    //     ch += 32;
    // printf("%c", ch);

    // int x, max;
    // for (int i = 1; i <= 6; i++)
    // {
    //     scanf("%d", &x);
    //     if (x > max)
    //         max = x;
    // }
    // printf("%d", max);

    // char a = 'f' - 'a';
    // printf("%d", a);

    // int a, b, c, *p = &c;
    // scanf("%d", p);
    // printf("%d%d", *p, c);

    // int k = 3, s[2];
    // s[0] = k;
    // k = s[1] * 10;
    // printf("%d", s[1]);

    // char a[] = "string";
    // int aaa(char *s);
    // printf("%d", aaa(a));

    // char ch= getchar();
    // while ((ch = getchar()) == 'b')   //getchar()取第一个字符  输入abcde   ch='a'
    //     printf("*");

    int a = 4, b = 3, c = 2, t = 0;
    if (b > a && a < c)
        t = a;
    a = c;
    c = t;
    printf("%d ", c);
}

int aaa(char *s)
{
    char *t = s;
    while (*t++) //先使t++指向下一个元素，在判断这个元素是否存在
        printf("%c\n", *t);
    t--;
    return (t - s);
}

int func(char ch, int n)
{
    printf("%c,%d\n", ch, n);
    return 1;
}

int fun(int a, int b)
{
    if (a > b)
        return (a);
    else
        return (b);
}
