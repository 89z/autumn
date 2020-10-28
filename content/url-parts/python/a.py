from urllib import parse

o = parse.ParseResult(
   fragment = '',
   netloc = 'example.com',
   params = '',
   path = '/one',
   query = 'two=even',
   scheme = 'https'
)

s = o.geturl()
print(s == 'https://example.com/one?two=even')
