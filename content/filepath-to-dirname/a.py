from os import path
s = 'C:\\Windows\\notepad.exe'
s1 = path.dirname(s)
print(s1 == 'C:\\Windows')
