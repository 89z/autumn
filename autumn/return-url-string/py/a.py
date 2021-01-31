from urllib import parse
s = 'one=odd&two=even'
m = parse.parse_qs(s)
print(m)
