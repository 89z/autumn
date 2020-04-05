import json
# example 1
r1 = open('a.json')
m1 = json.load(r1)
# example 2
r2 = open('a.json')
s1 = r2.read()
m2 = json.loads(s1)
# print
print(m1, m2)
