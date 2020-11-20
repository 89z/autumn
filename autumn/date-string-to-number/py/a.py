from datetime import datetime
o = datetime.strptime('2019-12-31', '%Y-%m-%d')
n = o.timestamp()
print(n == 1577772000)
