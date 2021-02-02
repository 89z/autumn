from urllib import parse
m = {'month': 'May', 'day': 'Friday'}
s = parse.urlencode(m)
print(s)
