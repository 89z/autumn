import re
s = 'Wednesday'
# example 1
o1 = re.search('^W', s)
# example 2
o2 = re.search('we', s, re.I)
# example 3
o3 = re.search('(?i)we', s)
# print
print(o1 != None, o2 != None, o3 != None)
