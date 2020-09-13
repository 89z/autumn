# example 1
a = ['May', 'June']
s = ','.join(a)
# example 2
a = [10, 11]
s2 = ','.join(map(str, a))
# print
print(s == 'May,June', s2 == '10,11')
