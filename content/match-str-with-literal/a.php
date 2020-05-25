<?php
$s = 'Sunday';
# example 1
$n = strpos($s, 'Su');
var_dump($n === 0);
# example 2
$n = strpos($s, 'un');
var_dump($n !== false);
# example 3
$n = stripos($s, 'su');
var_dump($n === 0);
# example 4
$s2 = $s[0];
var_dump($s2 == 'S');
# example 5
$s2 = $s[-1];
var_dump($s2 == 'y');
