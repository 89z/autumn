<?php
$m = ['year' => 2019, 'month' => 12];

while (true) {
   $s = key($m);
   if ($s === null) {
      break;
   }
   $n = current($m);
   echo $s, "\t", $n, "\n";
   next($m);
}
