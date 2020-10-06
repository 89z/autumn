m = {}
# example 1
m['year'] = 2019
# example 2
m2 = {'month': 12}
m.update(m2)
# example 3
m3 = {'day': 31}
m = {**m, **m3}
# print
print(m)
