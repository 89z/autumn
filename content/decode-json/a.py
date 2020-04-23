import json
# example 1
o1 = open('a.json')
m1 = json.load(o1)
# example 2
o2 = open('a.json')
s1 = o2.read()
m2 = json.loads(s1)
# print
print(m1, m2)
