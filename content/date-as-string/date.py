from datetime import date
o = date.today()
# example 1
s1 = o.isoformat()
# example 2
s2 = o.strftime('%Y-%m-%d')
# print
print(s1, s2)
