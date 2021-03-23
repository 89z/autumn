def number_format(d):
   e, f = d, 0
   while e >= 1e3:
      e /= 1e3
      f += 1
   return '%.3f' % e + ('', ' k', ' M', ' G')[f]

s = number_format(9012345678)
print(s == '9.012 G')
