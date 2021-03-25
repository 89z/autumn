from sys import stdout
from urllib import request
s = 'http://speedtest.lax.hivelocity.net'
b = request.urlopen(s).read()
stdout.buffer.write(b)
