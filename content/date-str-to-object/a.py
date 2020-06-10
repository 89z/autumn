# example 1
from datetime import datetime
o1 = datetime.strptime('2019-12-31', '%Y-%m-%d')
# example 2
import time
o2 = time.strptime('2019-12-31', '%Y-%m-%d')
# print
print(o1, o2, sep='\n')
