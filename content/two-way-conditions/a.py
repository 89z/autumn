n1 = 0
# example 1
if n1 > 0:
   s1 = 'positive'
elif n1 < 0:
   s1 = 'negative'
else:
   s1 = 'zero'
# example 2
s2 = 'negative' if n1 < 0 else 'positive'
# print
print(s1, s2)
