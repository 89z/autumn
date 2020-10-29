import pathlib
o = pathlib.Path(r'C:\Windows\notepad.exe')
s = str(o.parent)
print(s == r'C:\Windows')
