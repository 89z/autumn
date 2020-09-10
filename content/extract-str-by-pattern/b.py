import re
s = 'January'
# example 1
o = re.search('a(..)', s)
a = o.groups()
# example 2
a2 = re.findall('a(..)', s)
# print
print(a, a2)
