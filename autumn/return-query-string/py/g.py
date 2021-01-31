from urllib import parse

u = parse.SplitResult(
   'http', 'docs.python.org', '/library', 'month=May&day=Friday', fragment=''
)

s = parse.urlunsplit(u)
print(s)
