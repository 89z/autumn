import re
s = 'January'
# example 1
o = re.search('a.', s)
s1 = o.group()
# example 2
a2 = re.findall('a.', s)
# print
print(s1, a2)
