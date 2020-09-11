import json
m = {'📗/📕': 9}
# example 1
s = json.dumps(m)
# example 2
s2 = json.dumps(m, ensure_ascii=False)
# print
print(s, s2)
