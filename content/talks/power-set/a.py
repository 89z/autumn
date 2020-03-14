def f1(a_in):
   a_out = [[]]
   for v_in in a_in:
      a_out += [a_each + [v_in] for a_each in a_out]
   return a_out
a1 = [10, 11, 12]
a2 = f1(a1)
print(a2)
