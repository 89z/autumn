import pathlib
o = pathlib.Path('index.md')
s = o.resolve()
print(s)
