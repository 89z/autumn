# example 1
n1 = 10
while n1 < 20:
   print(n1)
   n1 += 1
# example 2
def f1(n1):
   if n1 > 19:
      return
   print(n1)
   f1(n1 + 1)
f1(10)
