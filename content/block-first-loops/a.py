# example 1
n1 = 10
while True:
   print(n1)
   if n1 == 19:
      break
   n1 += 1
# example 2
def f1(n1):
   print(n1)
   if n1 >= 19:
      return
   f1(n1 + 1)
f1(10)
