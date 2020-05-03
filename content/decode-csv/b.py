import csv
o_db = open('a.csv')
o_tab = csv.reader(o_db)
a_head = next(o_tab)
m_head = {}
for n_ind, s_col in enumerate(a_head):
   m_head[s_col] = n_ind
for a_row in o_tab:
   n_city = m_head['city']
   s_city = a_row[n_city]
   print(s_city)
