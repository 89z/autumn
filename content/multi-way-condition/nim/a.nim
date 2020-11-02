var
   n = 1 + 2
   s: string

case n
of 5:
   s = "five"
of 4, 3:
   s = "four or three"
else:
   s = $(n)

echo s == "four or three"
