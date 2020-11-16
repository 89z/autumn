def number_format(n):
   n2 = 0
   while n > 1024:
      n /= 1024
      n2 += 1
   return '%.3f' % n + ('', ' k', ' M', ' G')[n2]

s = number_format(1264)
print(s == '1.234 k')
