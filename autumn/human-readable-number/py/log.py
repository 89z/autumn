import math

def number_format(n):
   n2 = int(math.log10(n) / 3)
   return '%.3f' % (n / 1e3 ** n2) + ('', ' k', ' M', ' G')[n2]

s = number_format(9012345678)
print(s == '9.012 G')
