import json
in_s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}'
m = json.loads(in_s)
out_s = m['U2']['Boy'][0]
print(out_s == 'Twilight')
