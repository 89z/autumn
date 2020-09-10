import re
s = 'January'
# example 1
o = re.search('^J', s)
# example 2
o2 = re.search('ja', s, re.I)
# example 3
o3 = re.search('(?i)ja', s)
# print
print(o != None, o2 != None, o3 != None)
