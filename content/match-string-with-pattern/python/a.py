import re
s = 'January'
# example 1
o1 = re.search('^J', s)
# example 2
o2 = re.search('ja', s, re.I)
# example 3
o3 = re.search('(?i)ja', s)
# print
print(o1 != None, o2 != None, o3 != None)
