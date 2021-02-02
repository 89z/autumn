from urllib import parse

u = parse.SplitResult(
   'http', 'docs.python.org', '/library', 'month=May&day=Friday', fragment=''
)

s = u.geturl()
print(s)
