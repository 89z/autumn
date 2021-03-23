from datetime import datetime
import time
then = datetime.now()

while True:
   time.sleep(1 / 100)
   now = datetime.now() - then
   print(now, end='\r')
