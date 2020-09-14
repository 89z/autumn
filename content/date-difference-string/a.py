from datetime import datetime
o = datetime.fromisoformat('2019-12-31T00:00:00')
o2 = datetime.fromisoformat('2019-12-31T23:59:59')
o3 = o2 - o
print(o3)
