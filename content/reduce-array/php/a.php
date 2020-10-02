<?php
$a = ['May', 'June'];
# example 1
$s1 = array_reduce($a, fn ($s, $s1) => $s . $s1);
# example 2
$f = fn ($s, $s2) => $s . $s2;
$s2 = array_reduce($a, $f);
# print
var_dump($s1 == 'MayJune', $s2 == 'MayJune');
