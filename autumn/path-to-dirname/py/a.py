from os import path
s = r'C:\Windows\notepad.exe'
s = path.dirname(s)
print(s == r'C:\Windows')
