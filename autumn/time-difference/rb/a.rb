old_n = Time.now.to_f

loop do
   sleep(1 / 100)
   new_n = Time.now.to_f - old_n
   print "%.2f\r" % new_n
end
