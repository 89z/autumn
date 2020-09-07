import re
s = 'Wednesday'
# example 1
o = re.search('e.', s)
s = o.group()
# example 2
a = re.findall('e.', s)
# print
print(s, a)
