from os import path
s = 'C:\\Windows\\write.exe'
s2 = path.dirname(s)
print(s2 == 'C:\\Windows')
