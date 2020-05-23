from urllib import request
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
from sys import stdout
a1 = o1.read()
stdout.buffer.write(a1)
