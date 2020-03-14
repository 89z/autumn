a1 = [10, 11, 12]
P=[[]]
for i in a1:
   P += [s + [i] for s in P]
for s in P:
   print(s)
