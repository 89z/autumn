from datetime import datetime
s = '2020-12-31'
t = datetime.strptime(s, '%Y-%m-%d')
print(t)
