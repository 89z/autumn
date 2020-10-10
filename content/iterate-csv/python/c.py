import csv

def f(o_in):   
   a_in = csv.reader(o_in)
   a_out = []
   for n_row, a_row in enumerate(a_in):
      if n_row == 0:
         a_head = a_row
         continue
      m_row = dict(zip(a_head, a_row))
      a_out.append(m_row)
   return a_out   

o = open('a.csv')
a = f(o)
print(a)
