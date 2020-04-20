import re
s1 = '2019-12-31'
# example 1
o1 = re.search('\d', s1)
b1 = o1 != None
# example 2
o2 = re.search('-..', s1)
s2 = o2.group()
# example 3
a1 = re.findall('-..', s1)
# example 4
o3 = re.search('-(..)-(..)', s1)
a2 = o3.groups()
# print
print(b1, s2, a1, a2)
