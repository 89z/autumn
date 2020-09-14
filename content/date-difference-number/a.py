from datetime import datetime
o = datetime.strptime('2019-12-31', '%Y-%m-%d')
o2 = datetime.now()
n = (o2 - o).total_seconds()
print(n)
