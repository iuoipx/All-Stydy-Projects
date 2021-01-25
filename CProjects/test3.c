#include <stdio.h>
#include <string.h>
#define N 50
#define N1 2

// int fact(int m)
// {
//     int p;
//     if (m < 0)
//         printf("data error\n");
//     else if (m == 0 || m == 1)
//         p = 1; //正确定义结束条件，得
//     else
//         p = m * fact(m - 1); //正确写出递归表达式， 得 2 分。
//     return (p);              //正确返回结果，得 1
// }

struct Student
{
    char num[7], name[9];
    int score[3];
};

int main()
{
    // float x;
    // int a, b, c, i;
    // a = b = c = 0;
    // for (i = 1; i <= 10; i++)
    // {
    //     scanf("%f", &x);
    //     if (x > 0)
    //         a++;
    //     else if (x == 0)
    //         c++;
    //     else
    //         b++;
    // }
    // printf("%d %d %d", a, b, c);

    // int n;

    // int x;
    // scanf("%d", &n);
    // x = fact(n);
    // printf("n! = % ld", x);
    // void print(struct Student * stu);
    // struct Student stu[N1];
    // for (int i = 0; i < N1; i++)
    // {
    //     scanf("%s %s %d%d%d\n", &stu[i].num, &stu[i].name, &stu[i].score[0], &stu[i].score[1], &stu[i].score[2]);
    // }
    // print(stu);

    // char s[] = "hello\nworld";
    // printf("%s", s);

    // for (int i = 100; i <= 999; i++)
    // {
    //     int a, b, c;
    //     a = i % 10;
    //     b = i / 10 % 10;
    //     c = i / 100;
    //     if (a * a * a + b * b * b + c * c * c == i)
    //         printf("%d ", i);
    // }

    // static int x[5] = {1, 2, 3};
    // int y = *(x + 3) + 12;
    // printf("%d", y);
    // int a;
    // printf("%d", a == 2);

    // char str[][10] = {"China", "Beijing"}, *p = str;
    // printf("%s\n", p+10);

    //printf("%o", 010 << 1 ^ 1); //十进制17，输出八进制位21

    // int k;
    // float s;
    // for (k = 0, s = 0; k < 7; k++)
    //     s += k / 2;
    // printf("%d, %fn", k, s);

    // int a = 661, b = 888;
    // printf("%d\n", a, b); //661

    // int x = 11;
    // printf("%d", (x++ * 1 / 3));   //等价于x * 1 / 3

    // double d = 3.2;
    // int x, y;
    // x = 1.2;  //x此时变为1
    // y = (x + 3.8) / 5.0;
    // printf("%d \n", d * y);   //0

    // char *st = "how are you";
    // printf("%s", st + 2);

    // char arr[2][4];
    // strcpy(arr, "you");
    // strcpy(arr[1], "me");
    // arr[0][3] = '&'; //把arr[0]的最后一个元素由’\0’改为&
    // printf("%s \n", arr);

    int i, s = 0;
    for (i = 1; i < 10; i += 2)
    {
        s += i + 1;
        printf("% d\n", s);
    }

    return 0;
}

void print(struct Student *stu)
{
    for (int i = 0; i < N1; i++)
    {
        printf("%s %s %d %d %d\n", (*(stu + i)).num, (*(stu + i)).name, (*(stu + i)).score[0], (*(stu + i)).score[1], (*(stu + i)).score[2]);
    }
}
