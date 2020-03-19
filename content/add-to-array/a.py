# example 1
a1 = ['Sun']
a1.append('Mon')
# exmaple 2
a2 = ['Sun']
a2.extend(['Mon'])
# example 3
a3 = ['Sun']
a3 = [*a3, 'Mon']
# example 4
a4 = ['Sun']
a4 += ['Mon']
# print
print(a1, a2, a3, a4)
