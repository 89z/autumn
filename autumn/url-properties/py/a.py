from urllib import parse
u = parse.urlsplit('http://python.org/library?month=May')
# example 1
s1 = u.netloc
# example 2
s2 = u.path
# example 3
s3 = u.query
# print
print(s1, s2, s3)
