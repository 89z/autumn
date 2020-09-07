import re
s = 'Wednesday'
# example 1
o = re.search('e(..)', s)
a = o.groups()
# example 2
a2 = re.findall('e(..)', s)
# print
print(a, a2)
