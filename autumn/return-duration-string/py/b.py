from datetime import datetime
import time
old_o = datetime.now()

while True:
   time.sleep(1 / 100)
   new_o = datetime.now() - old_o
   print(new_o, end='\r')
