#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define prt(a, b) \
    if (!a)       \
    b = a

#define F(X, Y) X *Y

int main()
{
    // int i;
    // float x;
    // char s[10];
    // scanf("%3d%f%4s", &i, &x, s);
    // printf("%d %6.3f %s\n", i, x, s);  //81635ac76201

    void func(int x, int y, char *s1);
    struct date
    {
        int a, b;
        //char *s;
        char s[5];
    } arg;
    arg.a = 27;
    arg.b = 3;
    //arg.s = "abcd";
    strcpy(arg.s, "abcd");
    func(arg.a, arg.b, arg.s);
    printf("arg.a=%d,arg.b=%d,arg.s=%s\n", arg.a, arg.b, arg.s);

    // float y = 0.0, a[] = {2.0, 4.0, 6.0, 8.0, 10.0}, *p;
    // int i;
    // p = &a[1];
    // for (i = 0; i < 3; i++)
    //     y += *(p + i);
    // printf("%f", y);

    // float s = 2.734;
    // printf("%1.2f\n", s);

    // int x = 1, y = 0;
    // if (x > y)
    //     prt(x, y);
    // else
    //     prt(y, x);
    // printf("% d,%d", x, y);

    // int b = 2;
    // printf("%d", (b << 1) & 5);

    // int a = 3, b = 4;
    // printf("%d\n", F(a + b, a - b));

    // char str[13], str1[11];
    // scanf("%s%s", str, str1);   //使用scanf输出字符串时注意把空格当成不同字符之间的分隔符
    // printf("%s,%s", str, str1);

    //printf("%d", 1, 2, 3);  //输出1
    //printf("%d", (1, 2, 3));  //输出3

    // int a = 10;
    // printf("%d,%o,%x", a, a, a); //10,12,a

    // int i, x, y;
    // i = x = y = 0;
    // do
    // {
    //     ++i;
    //     if (i % 2 != 0)
    //     {
    //         x = x + i;
    //         i++;
    //     }
    //     y = y + i++;
    // } while (i <= 7);
    // printf("%d,%d", x, y);

    // int a = 4, b = 3, c = 2, t = 0;
    // if (b > a && a < c)
    //     t = a;
    // a = c;
    // c = t;
    // printf("%d", c);

    // int b = 2;
    // int c = (b << 1) & 5;
    // printf("%d", c);

    // int a, b, c, d;
    // a = b = 5;
    // /*
    //     1.(a++)使用5运算,然后a=6;
    //     2.(a--)使用6运算,然后a=5;
    //     3.(a++)使用5运算,然后a=6
    // */
    // c = (a++) + (a--) - (a++);
    // /*
    //     1.(b++)使用5运算,然后b=6;
    //     2.(++b)使用7运算,然后b=7;
    //     3.(--a)使用5运算,然后a=5
    // */
    // d = (b++) - (++b) + (--a);
    // printf("%d %d %d %d", a, b, c, d);

    // int a = 2147483647;
    // printf("%d", a + 2);

    // char *str = "0\n01234\11456";
    // printf("%d", strlen(str));

    // char *name[] = {"Follow me", "BASIC", "Great Wall", "FORTRAN", "Computer design"};
    // char **p;
    // int i;
    // for (i = 0; i < 5; i++)
    // {
    //     p = name + i;
    //     printf("%s\n", *p + i);
    // }

    // struct Student
    // {
    //     int num;
    //     float score;
    //     struct Student *next;
    // };
    // struct Student a, b, c, *head, *p;
    // a.num = 1000;
    // a.score = 89.5;
    // b.num = 1001;
    // b.score = 90;
    // c.num = 1002;
    // c.score = 91;
    // head = &a;
    // a.next = &b;
    // b.next = &c;
    // c.next = NULL;
    // p = head;
    // do
    // {
    //     printf("%ld %.1f\n", p->num, p->score);
    //     p = p->next;
    // } while (p != NULL);

    // void check(int *);
    // int *pl, i;
    // pl = (int *)malloc(5 * sizeof(int));
    // for (i = 0; i < 5; i++)
    // {
    //     scanf("%d", pl + i);
    // }
    // check(pl);
    // return 0;

    // union un {
    //     int I;
    //     double b;
    //     char c[11];
    // };
    // struct byte
    // {
    //     int a[3];
    //     float b;
    //     union un c[2];
    // } r;
    // printf("% d", sizeof(r));  //48

    // int a[3][2]={1,2,3,4,5,6},(*p)[2]=a;
    // printf("%d",*(*(p+1)+1));

    // char ch[3][5] = {"AAAA", "BBBBB", "CC"};
    // printf("\"%s\"\n", ch[1]);

    // int x = 102212, y = 012;
    // printf("%2d, %2d\n", x, y);

    // char a[1000];
    // for (int i = 0; i < 1000; i++)
    //     a[i] = -1 - i;
    // printf("%d", strlen(a));  //char 范围255

    union bt {
        int k;
        char c[2];
    } a;
    a.k = 7;
    printf("%o,%o", a.c[0], a.c[1]);

    printf(main);
}

void check(int *p)
{
    int i;
    printf("They are fail:");
    for (i = 0; i < 5; i++)
        if (p[i] < 60)
            printf("%d ", p[i]);
    printf("!\n");
}

void func(int x, int y, char *s1)
{
    x -= 5;
    y += 10;
    *s1 = 'A';
    *(s1 + 1) = 'B';
}