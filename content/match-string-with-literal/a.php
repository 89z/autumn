<?php
$s = 'June';
# example 1
$n = strpos($s, 'Ju');
var_dump($n === 0);
# example 2
$n = strpos($s, 'un');
var_dump($n !== false);
# example 3
$n = stripos($s, 'ju');
var_dump($n === 0);
# example 4
$s4 = $s[0];
var_dump($s4 == 'J');
# example 5
$s5 = $s[-1];
var_dump($s5 == 'e');
