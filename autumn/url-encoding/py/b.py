from urllib import parse
s = 'http://docs.python.org/library?west=left'
u = parse.urlsplit(s)
print(u)
