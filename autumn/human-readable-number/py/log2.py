import math

def f(n2):
   n = math.log2(n2) / 10
   n = math.floor(n)
   return '%.3f' % (n2 / 1024 ** n) + ('', ' k', ' M', ' G')[n]

s = f(1264)
print(s)
