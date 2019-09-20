#include <stdio.h>
#include <cs50.h>
#include <math.h>

int main()
{
int c=0;
float change;
do
{
  change=get_float("Cash Owed: ");
}
while (change<0);

change*=100;
change=round(change);

while(change>0)
{
    if(change>=25)
    {
        change-=25;
        c+=1;
    }
    else if(change>=10)
    {
        change-=10;
        c+=1;
    }
    else if(change>=5)
    {
        change-=5;
        c+=1;
    }
    else
    {
        change-=1;
        c+=1;
    }
}

printf("Total coins: %d\n", c);
}
