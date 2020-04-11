# example 1
import os.path as o1
s1 = o1.abspath('.')
# example 2
import os
s2 = os.getcwd()
# print
print(s1, s2, sep='\n')
