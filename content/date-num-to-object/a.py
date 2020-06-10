n = 365 * 24 * 60 * 60
# example 1
from datetime import datetime
o1 = datetime.fromtimestamp(n)
# exmaple 2
from datetime import date
o2 = date.fromtimestamp(n)
# print
print(o1, o2)
