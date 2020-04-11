from urllib import request
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
s1 = o1.read().decode()
print(s1)
