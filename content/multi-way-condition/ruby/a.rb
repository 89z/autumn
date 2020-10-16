s = 'minute'

case s
when 'hour'
   n = 23
when 'minute', 'second'
   n = 59
else
   n = -1
end

puts n == 59
