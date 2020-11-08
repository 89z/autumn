import pathlib
o = pathlib.Path(r'C:\Windows\notepad.exe')
s = o.stem
print(s == 'notepad')
