# example 1
import time
n1 = time.time()
n2 = int(n1)
# example 2
from datetime import datetime
o1 = datetime.now()
n3 = o1.timestamp()
n4 = int(n3)
# print
print(n2, n4)
