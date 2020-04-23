import csv
o1 = open('a.csv')
o2 = csv.DictReader(o1)
for m1 in o2:
   s1 = m1['city']
   print(s1)
