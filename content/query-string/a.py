from urllib import parse
# example A
s = 'one=odd&two=even'
mA = parse.parse_qs(s)
# example B
m = {'one': 'odd', 'two': 'even'}
sB = parse.urlencode(m)
# print
print(mA, sB, sep='\n')
