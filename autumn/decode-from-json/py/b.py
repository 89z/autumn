import json
s = '{"month": 12, "day": 31}'
m = json.loads(s)
n = m['day']
print(n == 31)
