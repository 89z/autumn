echo "example 1"
block:
   var n = 10
   while n < 20:
      echo n
      n += 1

echo "example 2"
block:
   var n = 10
   while true:
      if n > 19:
         break
      echo n
      n += 1
