from urllib import parse
u = parse.urlsplit('http://python.org/library?west=left')

# example 1
print(u.netloc == 'python.org')

# example 2
print(u.path == '/library')

# example 3
print(u.query == 'west=left')
