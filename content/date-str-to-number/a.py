from datetime import datetime
s = '2019-12-31'
o = datetime.strptime(s, '%Y-%m-%d')
n = o.timestamp()
print(n)
