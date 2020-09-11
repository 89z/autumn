a = [10, 11]
# example 1
a.append(13)
# example 2
a2 = [14, 15]
a.extend(a2)
# example 3
a = [*a, 16]
# example 4
a += [17, 18]
# print
print(a)
