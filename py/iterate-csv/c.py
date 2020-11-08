import csv
o_file = open('a.csv')
o_csv = csv.reader(o_file)
a_head = next(o_csv)

for a_row in o_csv:
   m = dict(zip(a_head, a_row))
   print(m)
