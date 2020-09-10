import json
m = {'📗/📕': 10}
# example 1
s1 = json.dumps(m)
# example 2
s2 = json.dumps(m, ensure_ascii=False)
# print
print(s1, s2)
