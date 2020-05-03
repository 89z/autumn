from urllib import parse
# example 1
s1 = 'https://example.com/one?two=even'
o1 = parse.urlparse(s1)
# example 2
o2 = parse.ParseResult(
   fragment = '',
   netloc = 'example.com',
   params = '',
   path = '/one',
   query = 'two=even',
   scheme = 'https'
)
s2 = o1.geturl()
# print
print(o1, s2)
