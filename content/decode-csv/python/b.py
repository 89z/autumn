import csv

def f(o_in):
   a_out = []
   for m_row in csv.DictReader(o_in):
      a_out.append(m_row)
   return a_out

o = open('a.csv')
a = f(o)
print(a)
