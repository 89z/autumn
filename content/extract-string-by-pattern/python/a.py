import re
s = 'January'
# example 1
a1 = re.findall('a.', s)
# example 2
a2 = re.findall('a(..)', s)
# print
print(a1, a2)
