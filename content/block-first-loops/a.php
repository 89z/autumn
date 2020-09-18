<?php

# example 1
$n = 10;
do {
   var_dump($n);
   $n++;
} while($n < 20);

# example 2
$n = 10;
while (true) {
   var_dump($n);
   if ($n == 19) {
      break;
   }
   $n++;
}

# example 3
$n = 10;
for (;;) {
   var_dump($n);
   if ($n == 19) {
      break;
   }
   $n++;
}
