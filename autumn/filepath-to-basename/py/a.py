from os import path
s = r'C:\Windows\notepad.exe'
s2 = path.basename(s)
print(s2 == 'notepad.exe')
