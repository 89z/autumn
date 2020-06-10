from datetime import date
n = 365 * 24 * 60 * 60
o = date.fromtimestamp(n)
# example 1
s1 = o.isoformat()
# example 2
s2 = o.strftime('%F')
# print
print(s1, s2)
