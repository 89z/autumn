def number_format(n)
   n2 = Math.log10(n).to_i / 3
   return '%.3f' % (n / 1e3 ** n2) + ['', ' k', ' M', ' G'][n2]
end

s = number_format(9012345678)
puts s == '9.012 G'
