import math

def number_format(d):
   e = int(math.log10(d) / 3)
   return '%.3f' % (d / 1e3 ** e) + ('', ' k', ' M', ' G')[e]

s = number_format(9012345678)
print(s == '9.012 G')
