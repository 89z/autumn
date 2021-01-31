from urllib import parse
s = 'http://docs.python.org/library?month=May&day=Friday'
u = parse.urlparse(s)
print(u)
