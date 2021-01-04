from datetime import datetime
s = '2020-05-04'
o = datetime.strptime(s, '%Y-%m-%d')
print(o)
