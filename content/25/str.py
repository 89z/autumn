s = '2.9'
# example 1
n = float(s)
# example 2
import locale
n2 = locale.atof(s)
# print
print(n == 2.9, n2 == 2.9)
