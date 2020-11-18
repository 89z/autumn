def number_format(n)
   n2, n3 = n, 0
   while n2 >= 1000
      n2 /= 1000.0
      n3 += 1
   end
   return '%.3f' % n2 + ['', ' k', ' M', ' G'][n3]
end

s = number_format(9012345678)
puts s == '9.012 G'
