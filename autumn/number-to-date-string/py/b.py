from datetime import date
n = 1577858399
o = date.fromtimestamp(n)
s = o.strftime('%a %b %#d %Y')
print(s == 'Tue Dec 31 2019')
