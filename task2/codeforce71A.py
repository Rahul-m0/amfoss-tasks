p=int(input())
for t in range(0,p):
    y=str(input())
    k=len(y)
    if k<=10:
        print(y)
    else:
        print(y[0]+str(k-2)+y[k-1])
