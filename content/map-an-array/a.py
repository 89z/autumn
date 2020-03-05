import math
a1 = [16, 25]
# example 1
a2 = [math.sqrt(n1) for n1 in a1]
# example 2
a3 = list(map(math.sqrt, a1))
# print
print(a2, a3)
