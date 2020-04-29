from urllib import parse
s1 = 'https://example.com/one?two=even'
o1 = parse.urlparse(s1)
s2 = o1.geturl()
print(s2)
