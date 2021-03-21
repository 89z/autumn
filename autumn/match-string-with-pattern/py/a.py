import re
s = 'January'

# example 1
r = re.search('ja', s, re.I)
print(r != None)

# example 2
r = re.search('(?i)ja', s)
print(r != None)
