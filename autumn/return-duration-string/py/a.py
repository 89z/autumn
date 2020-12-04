import time
old_n = time.time()

while True:
   time.sleep(1 / 100)
   new_n = time.time() - old_n
   print('%.2f' % new_n, end='\r')
