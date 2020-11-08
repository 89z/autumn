import os
# example 1
a = os.listdir('.')
print(a)
# example 2
for root, dirs, files in os.walk('.'):
   print(root, dirs, files)
