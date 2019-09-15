#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the compareTriplets function below.
def compareTriplets(a, b):
    k=[0,0]
    for x in range (0,3):
        if a[x]>b[x]:
            k[0]+=1        
        elif a[x]<b[x]:
            k[1]+=1 
    return k
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a = list(map(int, input().rstrip().split()))

    b = list(map(int, input().rstrip().split()))

    result = compareTriplets(a, b)
