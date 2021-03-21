from urllib import parse
u = parse.urlsplit('http://python.org/library?month=May')
# example 1
n = u.netloc
# example 2
p = u.path
# example 3
q = u.query
# print
print(n == 'python.org', p == '/library', q == 'month=May')
