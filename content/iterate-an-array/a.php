<?php
$a1 = ['Sun', 'Mon'];
# example 1
foreach ($a1 as $s1) {
   var_dump($s1);
}
# example 2
for ($n1 = 0; $n1 < count($a1); $n1++) {
   $s1 = $a1[$n1];
   var_dump($s1);
}
# example 3
for (reset($a1); key($a1) !== null; next($a1)) {
   $s1 = current($a1);
   var_dump($s1);
}
# example 4
array_walk($a1, fn($s1) => var_dump($s1));
