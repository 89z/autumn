n = 1000
# example 1
s = format(n)
# example 2
s2 = str(n)
# example 3
s3 = '{:,}'.format(n)
# example 4
s4 = format(n, ',')
# print
print(s == '1000', s2 == '1000', s3 == '1,000', s4 == '1,000')
