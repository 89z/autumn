import re
s1 = 'package'
# example 1
s2 = re.search('p.', s1).group()
# example 2
a1 = re.findall('a.', s1)
# example 3
a2 = re.search('p(..)', s1).groups()
# example 4
a3 = re.findall('p(..)', s1)
# print
print(s2, a1, a2, a3)
