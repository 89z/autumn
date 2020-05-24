import re
s1 = 'Wednesday'
s2 = 'e(..)'
# example 1
o1 = re.search(s2, s1)
a1 = o1.groups()
# example 2
a2 = re.findall(s2, s1)
# print
print(a1, a2)
