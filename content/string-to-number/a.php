<?php
# example 1
$n1 = floatval('1.2a');
# example 2
$n2 = intval('12a');
# example 3
$n3 = (int) '12a';
# example 4
$n4 = + '012';
# print
var_dump($n1, $n2, $n3, $n4);
