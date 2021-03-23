import json
a = ['/', '📗']

# example 1
s = json.dumps(a)
print(s)

# example 2
s = json.dumps(a, ensure_ascii=False)
print(s)

# example 3
s = json.dumps(a, indent=1)
print(s)
