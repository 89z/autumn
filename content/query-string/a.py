import urllib.parse as o1
# example 1
s1 = 'one=odd&two=even'
m1 = o1.parse_qs(s1)
# example 2
m2 = {'one': 'odd', 'two': 'even'}
s2 = o1.urlencode(m2)
# print
print(m1, s2)
