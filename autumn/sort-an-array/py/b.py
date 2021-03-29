# example 1
a = [10, 9]
a.sort()
print(a)

# example 2
a = [
   {'month': 10, 'day': 31}, {'month': 11, 'day': 30}
]

f = lambda m: m['day']
a.sort(key=f)
print(a)
