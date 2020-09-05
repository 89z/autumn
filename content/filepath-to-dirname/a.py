from os import path
s1 = 'C:\\Windows\\write.exe'
s2 = path.dirname(s1)
print(s2 == 'C:\\Windows')
