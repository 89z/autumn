from urllib import request
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
a1 = o1.read()
print(a1)
