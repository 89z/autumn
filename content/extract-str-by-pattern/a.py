import re
s1 = 'Wednesday'
# example 1
s2 = re.search('e.', s1).group()
# example 2
a1 = re.findall('e.', s1)
# example 3
a2 = re.search('e(..)', s1).groups()
# example 4
a3 = re.findall('e(..)', s1)
# print
print(s2, a1, a2, a3)
