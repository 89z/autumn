from datetime import datetime
import time
old = datetime.now()

while True:
   time.sleep(1 / 100)
   new = datetime.now() - old
   print(new, end='\r')
