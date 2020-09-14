import re
s = 'January'
# example 1
o = re.search('a.', s)
s = o.group()
# example 2
a = re.findall('a.', s)
# print
print(s, a)
