import json
# example 1
s = '{"year": 2019, "month": 12}'
m1 = json.loads(s)
# example 2
o = open('a.json')
m2 = json.load(o)
# print
print(m1, m2, sep='\n')
