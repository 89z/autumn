import csv
r1 = open('a.csv')
r2 = csv.DictReader(r1)
for m1 in r2:
   print(m1['city'])
