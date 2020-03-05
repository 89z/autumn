# example 1
t1 = type(10)
# example 2
t2 = type(type(10))
# example 3
t3 = type(object())
# example 4
class Day: Sunday = 10
t4 = type(Day)
# print
print(t1, t2, t3, t4)
