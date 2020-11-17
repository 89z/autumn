def number_format(n):
   n2 = 0
   while n > 1000:
      n /= 1000
      n2 += 1
   return '%.3f' % n + ('', ' k', ' M', ' G')[n2]

s = number_format(9012345678)
print(s == '9.012 G')
