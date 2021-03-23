import json
o = open('a.json')
m = json.load(o)
s = m['U2']['Boy'][0]
print(s == 'Twilight')
