#include <stdio.h>
#include <string.h>

int main()
{
    char string[80], c;
    int word = 0, num = 0, i;
    gets(string);
    for (i = 0; (c = string[i]) != '\0'; i++)
        if (c == ' ')
            word = 0;
        else if (word == 0)
        {
            word = 1;
            num++;
        }
    printf("%d", num);
    return 0;
}