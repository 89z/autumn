from urllib import parse
s = 'west=left&east=right'
m = parse.parse_qs(s)
print(m)
