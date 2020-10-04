from os import path
s = r'C:\Windows\notepad.exe'
s1 = path.dirname(s)
print(s1 == r'C:\Windows')
