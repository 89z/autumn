<?php
$s1 = '10';
$n1 = 1.9;
# example 1
$n2 = (int)($s1);
# example 2
$n3 = (int)($n1);
# example 3
$n4 = intval($s1);
# example 4
$n5 = intval($n1);
# example 5
$n6 = +($s1);
# print
var_dump($n2, $n3, $n4, $n5, $n6);
