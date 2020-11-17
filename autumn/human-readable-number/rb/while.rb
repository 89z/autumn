def number_format(n)
   n2 = 0
   while n > 1000
      n /= 1000.0
      n2 += 1
   end
   return '%.3f' % n + ['', ' k', ' M', ' G'][n2]
end

s = number_format(9012345678)
puts s == '9.012 G'
