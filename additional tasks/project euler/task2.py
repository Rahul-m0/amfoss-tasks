#!/bin/python3

import sys
def marco(N):
    a = 1
    b = 2
    c = 0
    k = 0
    while c<=N:
        
        if a%2==0:
            k+=a 
        c = a+b
        a = b
        b = c
        
    print(str(k))

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    marco(n)
