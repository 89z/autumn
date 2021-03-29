from urllib import parse

u = parse.SplitResult(
   'http', 'docs.python.org', '/library', 'west=left', fragment=''
)

s = parse.urlunsplit(u)
print(s)
