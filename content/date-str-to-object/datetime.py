from datetime import datetime
s = '2019-12-31'
# example 1
o1 = datetime.strptime(s, '%Y-%m-%d')
# example 2
o2 = datetime.fromisoformat(s)
# print
print(o1, o2)
