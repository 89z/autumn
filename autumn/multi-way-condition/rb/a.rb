n = 0x63 - 0x20

case n
when 0x41
   s = 'A'
when 0x42, 0x62
   s = 'B'
else
   s = n.chr
end

puts s == 'C'
