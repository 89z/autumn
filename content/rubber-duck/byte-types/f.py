from urllib import request
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
a1 = o1.read()
import os, sys
os.write(sys.stdout.fileno(), a1)
