from datetime import datetime
o = datetime.strptime('2019-12-31', '%Y-%m-%d')
o1 = datetime.now()
n = (o1 - o).total_seconds()
print(n)
