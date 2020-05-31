# example 1
n = 10
while True:
   print(n)
   if n == 19:
      break
   n += 1
# example 2
def f(n):
   print(n)
   if n < 19:
      f(n + 1)
f(10)
