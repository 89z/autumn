<?php
$s1 = '1.9';
$n1 = 10;
# example 1
$n2 = (float)($s1);
# example 2
$n3 = (float)($n1);
# example 3
$n4 = +($s1);
# example 4
$n5 = floatval($s1);
# example 5
$n6 = floatval($n1);
# print
var_dump($n2, $n3, $n4, $n5, $n6);