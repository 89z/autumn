<?php
$m = ['year' => 2019, 'month' => 12];

echo "example 1\n";
while (true) {
   $n = current($m);
   if ($n === false) {
      break;
   }
   $s = key($m);
   echo $s, ':', $n, "\n";
   next($m);
}

echo "example 2\n";
foreach ($m as $n) {
   echo $n, "\n";
}

echo "example 3\n";
foreach ($m as $s => $n) {
   echo $s, ':', $n, "\n";
}
