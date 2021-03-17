import json
a = ['/', '📗']
# example 1
s1 = json.dumps(a)
# example 2
s2 = json.dumps(a, ensure_ascii=False)
# example 3
s3 = json.dumps(a, indent=1)
# print
print(s1, s2, s3, sep='\n')
