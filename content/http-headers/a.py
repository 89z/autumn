import urllib.request as v1
s1 = 'http://speedtest.lax.hivelocity.net/100mb.file'
v2 = v1.Request(s1, method='HEAD')
# example 1
o1 = v1.urlopen(v2).info()
# example 2
m1 = dict(o1)
