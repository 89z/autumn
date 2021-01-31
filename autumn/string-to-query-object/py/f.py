from urllib import parse

u = parse.ParseResult(
   'http',
   'docs.python.org',
   '/library',
   params='',
   query='month=May&day=Friday',
   fragment=''
)

s = parse.urlunparse(u)
print(s)
