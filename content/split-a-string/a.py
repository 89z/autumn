s1 = 'sun mon tue'
s2 = '''sun mon tue
wed thu fri
'''
# example 1
a1 = s1.split(' ')
# example 2
a2 = s1.split(maxsplit = 1)
# example 3
a3 = s1.split()
# example 4
a4 = s2.splitlines()
# print
print(a1, a2, a3, a4, sep='\n')
