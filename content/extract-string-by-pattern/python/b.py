import re
s = 'January'
# example 1
o = re.search('a.', s)
s1 = o.group()
# example 2
o = re.search('a(..)', s)
a2 = o.groups()
# print
print(s1, a2)
