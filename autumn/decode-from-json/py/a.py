import json
f = open('a.json')
m = json.load(f)
n = m['day']
print(n == 31)
