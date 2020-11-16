def f(n)
   n2 = 0
   while n > 1024
      n /= 1024.0
      n2 += 1
   end
   return '%.3f' % n + ['', ' k', ' M', ' G'][n2]
end

puts f(1264)
