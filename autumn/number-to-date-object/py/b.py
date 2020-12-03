from datetime import date
o = date.fromtimestamp(1577858399)
s = o.strftime('%Y-%m-%d')
print(s == '2019-12-31')
