import re
s = 'south north'

# example 1
t = re.search('o..', s).group()
print(t == 'out')

# example 2
a = re.search('o(..)', s).groups()
print(a[0] == 'ut')
