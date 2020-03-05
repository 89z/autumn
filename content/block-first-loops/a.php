<?php
# example 1
$n1 = 10;
do {
   var_dump($n1);
   $n1++;
} while($n1 < 20);
# example 2
$n2 = 10;
while (true) {
   var_dump($n2);
   if ($n2 == 19) {
      break;
   }
   $n2++;
}
# example 3
$n3 = 10;
for (;;) {
   var_dump($n3);
   if ($n3 == 19) {
      break;
   }
   $n3++;
}
