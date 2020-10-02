<?php

echo "example 1\n";
for ($n = 10; $n < 20; $n++) {
   var_dump($n);
}

echo "example 2\n";
$n = 10;
while ($n < 20) {
   var_dump($n);
   $n++;
}

echo "example 3\n";
$a = range(10, 19);
foreach ($a as $n) {
   var_dump($n);
}
