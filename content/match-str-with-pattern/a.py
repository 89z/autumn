import re
s1 = 'Sunday'
# example 1
o1 = re.search('^Su', s1)
# example 2
o2 = re.search('un.', s1)
# example 3
o3 = re.search('ay$', s1)
# example 4
o4 = re.search('su', s1, re.I)
# example 5
o5 = re.search('(?i)su', s1)
# print
print(o1 != None, o2 != None, o3 != None, o4 != None, o5 != None)
