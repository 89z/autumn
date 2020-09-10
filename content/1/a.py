import csv
o_db = open('a.csv')
for m_row in csv.DictReader(o_db):
   print(m_row)
