import csv
file_o = open('a.csv')
csv_o = csv.reader(file_o)
head_a = next(csv_o)

for row_a in csv_o:
   m = dict(zip(head_a, row_a))
   print(m)
