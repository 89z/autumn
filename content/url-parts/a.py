from urllib import parse
# example 1
s1 = 'https://example.com/one?two=even'
o1 = parse.urlparse(s1)
print(o1)
# example 2
s2 = o1.geturl()
print(s2)
