from urllib import parse

# example A
oA = parse.urlparse('https://example.com/one?two=even')

# example B
o = parse.ParseResult(
   fragment = '',
   netloc = 'example.com',
   params = '',
   path = '/one',
   query = 'two=even',
   scheme = 'https'
)
s = o.geturl()

# print
print(oA, s, sep='\n')
