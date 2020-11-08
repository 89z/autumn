import re
s = 'January'
# example 1
o1 = re.search('ja', s, re.I)
# example 2
o2 = re.search('(?i)ja', s)
# print
print(o1 != None, o2 != None)
