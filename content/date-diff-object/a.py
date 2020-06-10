from datetime import datetime
o1 = datetime.strptime('2019-12-31', '%Y-%m-%d')
o2 = datetime.now()
o3 = o2 - o1
n = o3.total_seconds()
print(n)
