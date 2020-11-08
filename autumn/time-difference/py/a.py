from datetime import datetime
o = datetime.fromisoformat('2019-12-31')
o1 = datetime.fromisoformat('2019-12-31T23:59:59')
s = str(o1 - o)
print(s == '23:59:59')
