import csv
o_db = open('a.csv')
o_tab = csv.reader(o_db)
a_head = next(o_tab)

for a_row in o_tab:
   m_row = dict(zip(a_head, a_row))
   s_city = m_row['city']
   print(s_city)
