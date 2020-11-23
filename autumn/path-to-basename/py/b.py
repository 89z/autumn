import pathlib
o = pathlib.Path(r'C:\Windows\notepad.exe')
s = o.name
print(s == 'notepad.exe')
