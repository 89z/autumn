<?php
$a = ['May', 'June'];

# example 1
$s = array_reduce($a, fn ($s, $t) => $s . $t);
var_dump($s == 'MayJune');

# example 2
$n = array_reduce($a, fn ($n, $s) => $n + strlen($s));
var_dump($n == 7);
