import locale
locale.setlocale(locale.LC_ALL, '')
s1 = '{:,}'.format(n)
s2 = format(n, ',')
s3 = str(n)
s4 = repr(n)
