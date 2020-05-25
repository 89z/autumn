from urllib import request
s = 'http://speedtest.lax.hivelocity.net'
o = request.urlopen(s)
from sys import stdout
y = o.read()
stdout.buffer.write(y)
