from sys import stdout
from urllib import request
s = 'https://speedtest.lax.hivelocity.net'
y = request.urlopen(s).read()
stdout.buffer.write(y)
