import re

# example 1
s = re.search('..n', 'Sunday Monday').group()
print(s == 'Sun')

# example 2
a = re.search('(..)n', 'Sunday Monday').groups()
print(a[0] == 'Su')
