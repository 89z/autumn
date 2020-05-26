import json
# example 1
o = open('a.json')
m1 = json.load(o)
# example 2
s = '{"Sunday": 10}'
m2 = json.loads(s)
# print
print(m1, m2)
