from os import path
s = r'C:\Windows\notepad.exe'
s2 = path.dirname(s)
print(s2 == r'C:\Windows')
