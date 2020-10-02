from urllib import parse
# example 1
s = 'one=odd&two=even'
m1 = parse.parse_qs(s)
# example 2
m = {'one': 'odd', 'two': 'even'}
s2 = parse.urlencode(m)
# print
print(m1, s2, sep='\n')
