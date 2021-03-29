import re
s = 'South North'

# example 1
r = re.search('so', s, re.I)
print(r != None)

# example 2
r = re.search('(?i)so', s)
print(r != None)
