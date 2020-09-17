<?php
$a = ['May', 'June'];
# example 1
$s = array_reduce($a, fn ($sa, $sc) => $sa . $sc);
# example 2
$f = fn ($sa, $sc) => $sa . $sc;
$s2 = array_reduce($a, $f);
# print
var_dump($s == 'MayJune', $s2 == 'MayJune');
