from pathlib import PurePath
p = PurePath('south', 'north')
s = str(p)
print(s == r'south\north')
