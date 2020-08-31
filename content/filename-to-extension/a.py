import pathlib
s1 = 'index.md'
o1 = pathlib.PurePath(s1)
s2 = o1.suffix
print(s2 == '.md')
