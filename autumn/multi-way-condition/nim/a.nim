var
   c = '\x43'
   n: int

case c
of 'A':
   n = 0x41
of 'B', 'b':
   n = 0x42
else:
   n = int c

echo n == 0x43
