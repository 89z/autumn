# example 1
class Date:
   year = 2019
   month = 12
o1 = Date()

# example 2
import types
o2 = types.SimpleNamespace(year = 2019, month = 12)

# print
print(o1.year == 2019, o2.year == 2019)