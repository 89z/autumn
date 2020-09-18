n = 1000
# example 1
s = '{:,}'.format(n)
# example 2
s2 = format(n, ',')
# example 3
import locale
locale.setlocale(locale.LC_ALL, '')
s3 = format(n, 'n')
# print
print(s == '1,000', s2 == '1,000', s3 == '1,000')
