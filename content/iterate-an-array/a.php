<?php
$a = ['Sunday', 'Monday'];
# example 1
foreach ($a as $s) {
   var_dump($s);
}
# example 2
for ($n = 0; $n < count($a); $n++) {
   $s = $a[$n];
   var_dump($s);
}
# example 3
for (reset($a); key($a) !== null; next($a)) {
   $s = current($a);
   var_dump($s);
}
# example 4
array_walk($a, fn($s) => var_dump($s));
