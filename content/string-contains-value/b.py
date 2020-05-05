import re
s1 = 'Sunday'
# example 1
o1 = re.search('^Su', s1)
print(o1 != None)
# example 2
o2 = re.search('un.', s1)
print(o2 != None)
# example 3
o3 = re.search('ay$', s1)
print(o2 != None)
