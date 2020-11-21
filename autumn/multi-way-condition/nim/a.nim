var
   n = 0x63 - 0x20
   c: char

case n
of 0x41:
   c = 'A'
of 0x42, 0x62:
   c = 'B'
else:
   c = char n

echo c == 'C'
