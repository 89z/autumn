def number_format(d)
   e = Math.log10(d).to_i / 3
   return '%.3f' % (d / 1e3 ** e) + ['', ' k', ' M', ' G'][e]
end

s = number_format(9012345678)
puts s == '9.012 G'
