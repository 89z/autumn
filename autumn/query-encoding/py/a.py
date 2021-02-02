from urllib import parse
s = 'month=May&day=Friday'
m = parse.parse_qs(s)
print(m)
