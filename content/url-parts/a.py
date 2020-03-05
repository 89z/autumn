from urllib.parse import urlparse
s1 = 'http://sun.com/mon?tue=10'
m1 = urlparse(s1)
print(m1)
