import time
then = time.time()

while True:
   time.sleep(1 / 100)
   now = time.time() - then
   print('%.2f' % now, end='\r')
