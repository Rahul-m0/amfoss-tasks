import os
import sys



def timeConversion(s):
    coc=s.split(':')
    boo=coc[-1][2:4]
    r=coc[0]
    print(r)
    if boo=="AM":
        dic=coc[-1].replace("AM","")
        if coc[0]==12:
            r="00"
    elif boo=="PM":
        dic=coc[-1].replace("PM","") 
        if coc[0] != 12:
            r=int(coc[0])+12
    return str(str(r) + ":" + coc[1] + ":" + dic)
  

if __name__ == '__main__':
    f = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)
 
    f.write(result + '\n')

    f.close()
