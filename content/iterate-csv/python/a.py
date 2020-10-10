import csv

def f(s_in):
   a_in = s_in.splitlines()
   a_out = []
   for m_row in csv.DictReader(a_in):
      a_out.append(m_row)
   return a_out

s = '''Month,Day
January,Sunday
February,Monday'''

a = f(s)
print(a)
