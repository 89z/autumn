import csv
o = open('a.csv')

for m in csv.DictReader(o):
   print(m)
