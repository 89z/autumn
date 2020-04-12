from sys import stdout
from urllib import request
import os
s1 = 'http://speedtest.lax.hivelocity.net'
o1 = request.urlopen(s1)
a1 = o1.read()
# example 1
n1 = stdout.fileno()
os.write(n1, a1)
# example 2
os.write(1, a1)
