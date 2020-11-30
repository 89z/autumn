<?php
$m = ['month' => 12, 'day' => 31];

while (true) {
   $s = key($m);
   if ($s === null) {
      break;
   }
   $n = current($m);
   echo $s, "\t", $n, "\n";
   next($m);
}
