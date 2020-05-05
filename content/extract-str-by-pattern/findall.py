import re
s1 = '2019-12-31'
# example 1
a1 = re.findall('\d', s1)
b1 = len(a1) != 0
# example 2
a2 = re.findall('-..', s1)
s2 = a2[0]
# example 3
a3 = re.findall('-..', s1)
# example 4
a4 = re.findall('-(..)-(..)', s1)
a5 = a4[0]
# print
print(b1, s2, a3, a5)
