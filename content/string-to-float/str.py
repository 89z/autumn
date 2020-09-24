s = '9.9'
# example 1
n1 = float(s)
# example 2
import locale
n2 = locale.atof(s)
# print
print(n1 == 9.9, n2 == 9.9)
