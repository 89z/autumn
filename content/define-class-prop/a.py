# example 1
class Jan:
   Sun = 10
   Mon = 11
o1 = Jan()
# example 2
import types
o2 = types.SimpleNamespace(Sun = 10, Mon = 11)
# print
print(o1, o2)
