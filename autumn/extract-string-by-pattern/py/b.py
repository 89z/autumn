import re
s = 'Sunday Monday'
# example 1
s1 = re.search('..n', s).group()
print(s1 == 'Sun')
# example 2
a = re.search('(..)n', s).groups()
print(a[0] == 'Su')
