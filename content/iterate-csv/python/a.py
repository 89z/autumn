import csv

s = '''Month,Day
January,Sunday
February,Monday'''

a = s.splitlines()

for m in csv.DictReader(a):
   print(m)
