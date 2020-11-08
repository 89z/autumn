var
   n = 1 + 2
   b: bool

case n
of 5:
   b = false
of 4, 3:
   b = true
else:
   b = n < 3

echo b
