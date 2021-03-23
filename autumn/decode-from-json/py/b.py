import json
src = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}'
m = json.loads(src)
dst = m['U2']['Boy'][0]
print(dst == 'Twilight')
