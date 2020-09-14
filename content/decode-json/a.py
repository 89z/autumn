import json
# example 1
s = '{"year": 2019, "month": 12, "day": 31}'
m = json.loads(s)
# example 2
o = open('a.json')
m2 = json.load(o)
# print
print(m, m2, sep='\n')
