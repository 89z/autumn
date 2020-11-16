def f(n):
   n2 = 0
   while n > 1000:
      n /= 1000
      n2 += 1
   return '%.3f' % n + ('', ' k', ' M', ' G')[n2]

s = f(1234)
print(s)
