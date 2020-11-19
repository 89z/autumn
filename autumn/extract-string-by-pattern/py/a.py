import re
s = 'Sunday Monday'
# example 1
a = re.findall('..n', s)
print(a[0] == 'Sun', a[1] == 'Mon')
# example 2
a = re.findall('(..)n', s)
print(a[0] == 'Su', a[1] == 'Mo')
