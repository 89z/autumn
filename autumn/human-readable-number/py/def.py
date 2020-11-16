def f(n, n2 = 0):
   if n >= 1000:
      return f(n / 1024, n2 + 1)
   return '%.3f' % n + ('', ' k', ' M', ' G')[n2]

s = f(1264)
print(s)
