<?php
$m1 = ['Sun' => 10, 'Mon' => 11];
# example 1
foreach ($m1 as $n1) {
   var_dump($n1);
}
# example 2
while (true) {
   $n1 = current($m1);
   if ($n1 === false) {
      break;
   }
   $s1 = key($m1);
   var_dump($s1, $n1);
   next($m1);
}
