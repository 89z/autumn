# example 1
class Date:
   year = 2020
o = Date()

# example 2
import types
o2 = types.SimpleNamespace(year = 2020)

# print
print(o.year == 2020, o2.year == 2020)
