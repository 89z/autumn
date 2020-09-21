from os import path
s = 'C:\\Windows\\notepad.exe'
s1 = path.basename(s)
print(s1 == 'notepad.exe')
