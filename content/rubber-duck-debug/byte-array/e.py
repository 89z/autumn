from urllib import request
import sys
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
a1 = o1.read()
sys.stdout.buffer.write(a1)
