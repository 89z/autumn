from sys import stdout
from urllib import request
s = 'https://speedtest.lax.hivelocity.net'
o = request.urlopen(s)
y = o.read()
stdout.buffer.write(y)
