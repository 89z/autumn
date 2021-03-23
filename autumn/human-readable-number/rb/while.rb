def number_format(d)
   e, f = d, 0
   while e >= 1000
      e /= 1000.0
      f += 1
   end
   return '%.3f' % e + ['', ' k', ' M', ' G'][f]
end

s = number_format(9012345678)
puts s == '9.012 G'
