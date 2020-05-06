<?php
$s1 = 'Sunday';
# example 1
$n1 = preg_match('/^Su/', $s1);
var_dump($n1 == 1);
# example 2
$n2 = strpos($s1, 'Su');
var_dump($n2 === 0);
# example 3
$n3 = preg_match('/un./', $s1);
var_dump($n3 == 1);
# example 4
$n4 = strpos($s1, 'un');
var_dump($n4 !== false);
# example 5
$n5 = preg_match('/ay$/', $s1);
var_dump($n5 == 1);
