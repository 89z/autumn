from datetime import datetime
o1 = datetime.strptime('2019-12-31', '%Y-%m-%d')
o2 = datetime.now()
n = (o2 - o1).total_seconds()
print(n)
