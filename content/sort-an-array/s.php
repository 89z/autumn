<?php
$a = ['April', 'May', 'June'];
# example 1
sort($a);
print_r($a);
# example 2
$f = fn ($s, $s2) => strlen($s) <=> strlen($s2);
usort($a, $f);
print_r($a);
