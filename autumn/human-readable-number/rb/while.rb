def number_format(n)
   n2 = 0
   while n > 1024
      n /= 1024.0
      n2 += 1
   end
   return '%.3f' % n + ['', ' k', ' M', ' G'][n2]
end

s = number_format(1264)
puts s == '1.234 k'
