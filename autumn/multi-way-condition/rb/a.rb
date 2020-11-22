s = "\x43"

case s
when 'A'
   n = 0x41
when 'B', 'b'
   n = 0x42
else
   n = s.ord
end

puts n == 0x43
