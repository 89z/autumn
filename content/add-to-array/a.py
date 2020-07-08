# example 1
a1 = [10, 11]
a1.append(12)
# example 2
a2 = [10, 11]
a2.extend([12, 13])
# example 3
a3 = [10, 11]
a3 = [*a3, 12]
# example 4
a4 = [10, 11]
a4 += [12, 13]
# print
print(a1, a2, a3, a4)
