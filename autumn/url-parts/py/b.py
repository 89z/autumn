from urllib import parse
m = {'one': 'odd', 'two': 'even'}
s = parse.urlencode(m)
print(s)
