<?php
$a = ['BBBB', 'AA', 'CCC'];
# example 1
sort($a);
print_r($a);
# example 2
rsort($a);
print_r($a);
# example 3
$f = fn ($s, $s2) => strlen($s) <=> strlen($s2);
usort($a, $f);
print_r($a);
