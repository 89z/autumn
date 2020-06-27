n = 10
# example 1
if n > 12:
   s1 = 'Tue'
elif n > 11:
   s1 = 'Mon'
else:
   s1 = 'Sun'
# example 2
s2 = 'Mon' if n > 11 else 'Sun'
# print
print(s1, s2)
