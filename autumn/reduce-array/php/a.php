<?php
$a = ['May', 'June'];
# example 1
$s = array_reduce($a, fn ($s, $s1) => $s . $s1);
# example 2
$n = array_reduce($a, fn ($n, $s) => $n + strlen($s));
# print
var_dump($s == 'MayJune', $n == 7);
