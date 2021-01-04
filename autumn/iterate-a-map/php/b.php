<?php
$m = ['month' => 5, 'day' => 4];

while (true) {
   $s = key($m);
   if ($s === null) {
      break;
   }
   $n = current($m);
   echo $s, "\t", $n, "\n";
   next($m);
}
