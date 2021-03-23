then = Time.now.to_f

loop do
   sleep(1 / 100)
   now = Time.now.to_f - then
   print "%.2f\r" % now
end
