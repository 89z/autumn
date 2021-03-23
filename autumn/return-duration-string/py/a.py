import time
old = time.time()

while True:
   time.sleep(1 / 100)
   new = time.time() - old
   print('%.2f' % new, end='\r')
