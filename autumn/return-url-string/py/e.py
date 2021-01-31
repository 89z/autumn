from urllib import parse
s = 'http://docs.python.org/library?month=May&day=Friday'
u = parse.urlsplit(s)
print(u)
