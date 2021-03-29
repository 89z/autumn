import re
s = 'south north'
# example 1
a = re.findall('o..', s)
print(a[0] == 'out', a[1] == 'ort')
# example 2
a = re.findall('o(..)', s)
print(a[0] == 'ut', a[1] == 'rt')
