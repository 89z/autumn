import math

def number_format(n):
   n2 = math.log(n, 1024)
   n3 = math.floor(n2)
   return '%.3f' % (n / 1024 ** n3) + ('', ' k', ' M', ' G')[n3]

s = number_format(1264)
print(s == '1.234 k')
