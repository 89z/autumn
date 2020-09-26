<?php
$a = ['May', 'June'];

echo "example 1\n";
foreach ($a as $s) {
   echo $s, "\n";
}

echo "example 2\n";
foreach ($a as $n => $s) {
   echo $n, "\t", $s, "\n";
}

echo "example 3\n";
$n = count($a) - 1;
while ($n >= 0) {
   echo $n, "\t", $a[$n], "\n";
   $n--;
}

echo "example 4\n";
$n = array_key_last($a);
while ($n >= 0) {
   echo $n, "\t", $a[$n], "\n";
   $n--;
}

echo "example 5\n";
$f = fn ($s) => var_dump($s);
array_walk($a, $f);
