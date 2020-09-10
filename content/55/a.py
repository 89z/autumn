from urllib import parse
# example 1
s = 'one=odd&two=even'
m = parse.parse_qs(s)
# example 2
m2 = {'one': 'odd', 'two': 'even'}
s2 = parse.urlencode(m2)
# print
print(m, s2, sep='\n')
