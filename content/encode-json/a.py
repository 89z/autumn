import json
a = ['/', '📗']
# example 1
s1 = json.dumps(a)
# example 2
s2 = json.dumps(a, ensure_ascii=False)
# print
print(s1, s2, sep='\n')
