#include<stdlib.h>
#include<stdio.h>

int main()
{
    int i, j, h;


    printf("Enter height of mario level: ");
    scanf("%d", &h);


    for(i=2; i<=h+1; i++)
    {

        for(j=i; j<h+1; j++)
        {
            printf(" ");
        }

        for(j=1; j<=i; j++)
        {
            printf("#");
        }
        printf("\n");
    }
}
