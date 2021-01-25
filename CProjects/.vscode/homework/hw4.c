#include <stdio.h>

int main()
{
    int count = 0, i, j;
    for (i = 1; i <= 9; i++)
    {
        for (j = 1; j <= 9; j++)
        {
            printf("%2d ", j * i);
            if (j == 9)
                printf("\n");
        }
    }
    return 0;
}