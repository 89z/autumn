import os

for root, dirs, files in os.walk('..'):
   print(root, dirs, files)
