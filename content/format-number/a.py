n1 = 1000
# example 1
s1 = format(n1, '5')
# example 2
s2 = format(n1, ',')
# example 3
import locale
locale.setlocale(locale.LC_ALL, '')
s3 = format(n1, 'n')
# print
print([s1, s2, s3])
