# example 1
class jan:
   sun = 10
   mon = 11
o1 = jan()
# example 2
import types
o2 = types.SimpleNamespace(sun = 10, mon = 11)
# print
print(o1, o2)
