<?php
$s1 = 'Sun Mon';
# example 1
$n1 = strpos($s1, 'n');
# example 2
$n2 = strrpos($s1, 'n');
# example 3
$n3 = stripos($s1, 'N');
# example 4
$n4 = strripos($s1, 'N');
# print
var_dump($n1, $n2, $n3, $n4);
