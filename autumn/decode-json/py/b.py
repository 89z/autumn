import json
s = '{"month": 12, "day": 31}'
m = json.loads(s)
print(m['day'] == 31)
