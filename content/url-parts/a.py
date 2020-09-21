from urllib import parse

# example 1
o1 = parse.urlparse('https://example.com/one?two=even')

# example 2
o = parse.ParseResult(
   fragment = '',
   netloc = 'example.com',
   params = '',
   path = '/one',
   query = 'two=even',
   scheme = 'https'
)
s2 = o.geturl()

# print
print(o1, s2, sep='\n')
