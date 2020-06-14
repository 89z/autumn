<?php
$m = ['Sun' => 10, 'Mon' => 11];
# example 1
while (true) {
   $n = current($m);
   if ($n === false) {
      break;
   }
   $s = key($m);
   var_dump($s, $n);
   next($m);
}
# example 2
foreach ($m as $n) {
   var_dump($n);
}
# example 3
foreach ($m as $s => $n) {
   var_dump($s, $n);
}
