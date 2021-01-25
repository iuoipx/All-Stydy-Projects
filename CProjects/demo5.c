#include <stdio.h>

int main()
{
    void sort(int *p, int n);
    int i, *p, a[10];
    p = a;
    printf("please input ten integer numbers :");
    for (i = 0; i < 10; i++)
    {
        scanf("%d", p++);
    }
    printf("\n");
    p = a;
    sort(p, 10);
    for (i = 0; i < 10; i++)
    {
        printf("%d ", *p);
        p++;
    }
    return 0;
}

void sort(int *p, int n)
{
    int i, j, k, temp;
    for (i = 0; i < n - 1; i++)
    {
        k = i;
        for (j = i + 1; j < n; j++)
        {
            if (*(p + i) < *(p + j))
            {
                k = j;
                if (k != i)
                {
                    temp = *(p + i);
                    *(p + i) = *(p + k);
                    *(p + k) = temp;
                }
            }
        }
    }
}