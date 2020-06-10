from datetime import date
n = 366 * 24 * 60 * 60
o = date.fromtimestamp(n)
s = o.strftime('%a %b %#d %Y')
print(s == 'Fri Jan 1 1971')
