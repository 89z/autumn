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
for ($n = 0; $n < count($a); $n++) {
   echo $n, "\t", $a[$n], "\n";
}

echo "example 4\n";
for (reset($a); key($a) !== null; next($a)) {
   echo current($a), "\n";
}

echo "example 5\n";
$f = fn ($s) => var_dump($s);
array_walk($a, $f);
