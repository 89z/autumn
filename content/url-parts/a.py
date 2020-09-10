from urllib import parse

# example 1
o = parse.urlparse('https://example.com/one?two=even')

# example 2
o2 = parse.ParseResult(
   fragment = '',
   netloc = 'example.com',
   params = '',
   path = '/one',
   query = 'two=even',
   scheme = 'https'
)
s = o2.geturl()

# print
print(o, s, sep='\n')
