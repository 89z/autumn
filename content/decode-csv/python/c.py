import csv
s_in = '''Month,Day
January,Sunday
February,Monday'''
a_in = s_in.splitlines()
a_out = []
for m_row in csv.DictReader(a_in):
   print(m_row)
