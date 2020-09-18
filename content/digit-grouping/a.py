n = 1000
# example A
sA = '{:,}'.format(n)
# example B
sB = format(n, ',')
# example C
import locale
locale.setlocale(locale.LC_ALL, '')
sC = format(n, 'n')
# print
print(sA == '1,000', sB == '1,000', sC == '1,000')
