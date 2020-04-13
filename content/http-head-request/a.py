from urllib import request
s1 = 'http://speedtest.lax.hivelocity.net/100mb.file'
o1 = request.Request(s1, method='HEAD')
# example 1
o2 = request.urlopen(o1).info()
# example 2
m1 = dict(request.urlopen(o1).info())
# print
print(o2, m1)
