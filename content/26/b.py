import locale
locale.setlocale(locale.LC_ALL, '')
n = 1000
s = format(n, 'n')
print(s == '1,000')
