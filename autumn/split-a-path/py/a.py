from os import path
s = r'C:\Windows\notepad.exe'
s = path.basename(s)
print(s == 'notepad.exe')
