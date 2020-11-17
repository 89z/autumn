def number_format(n)
   n2 = Math.log10(n) / 3
   return '%.3f' % (n / 1000 ** n2) + ['', ' k', ' M', ' G'][n2]
end

s = number_format(9012345678)
puts s
