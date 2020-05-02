<?php
# example 1
for ($n1 = 10; $n1 < 20; $n1++) {
   var_dump($n1);
}
# example 2
$n1 = 10;
while ($n1 < 20) {
   var_dump($n1);
   $n1++;
}
# example 3
$a1 = range(10, 19);
foreach ($a1 as $n1) {
   var_dump($n1);
}
# example 4
function f1($n1) {
   if ($n1 > 19) {
      return;
   }
   var_dump($n1);
   f1($n1 + 1);
}
f1(10);
