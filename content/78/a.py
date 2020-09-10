# example 1
def f(s):
   return len(s)
# example 2
f2 = lambda s: len(s)
# print
n = f('May')
n2 = f2('May')
print(n == 3, n2 == 3)
