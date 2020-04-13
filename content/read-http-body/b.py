from urllib import parse
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = parse.urlparse(s1)
from http import client
o2 = client.HTTPConnection(o1.hostname)
o2.request('GET', o1.path)
o3 = o2.getresponse()
from sys import stdout
a1 = o3.read()
stdout.buffer.write(a1)
