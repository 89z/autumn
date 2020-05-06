<?php
$s1 = 'Sunday';
# example 1
$n1 = strpos($s1, 'Su');
var_dump($n1 === 0);
# example 2
$n2 = strpos($s1, 'un');
var_dump($n2 !== false);
# example 3
$n3 = stripos($s1, 'su');
var_dump($n3 === 0);
