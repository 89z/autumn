def number_format(n):
   n2, n3 = n, 0
   while n2 >= 1e3:
      n2 /= 1e3
      n3 += 1
   return '%.3f' % n2 + ('', ' k', ' M', ' G')[n3]

s = number_format(9012345678)
print(s == '9.012 G')
