import csv
o = open('a.csv')
o = csv.reader(o)
head_a = next(o)

for row_a in o:
   m = dict(zip(head_a, row_a))
   print(m)
