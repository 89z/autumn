<?php
$a = ['BBBB', 'AA', 'CCC'];
# example A
sort($a);
print_r($a);
# example B
rsort($a);
print_r($a);
# example C
$f = fn ($sA, $sB) => strlen($sA) <=> strlen($sB);
usort($a, $f);
print_r($a);
