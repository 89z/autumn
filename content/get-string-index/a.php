<?php
$s1 = 'sun mon tue';
# example 1
$n1 = strpos($s1, ' ');
# example 2
$n2 = strrpos($s1, ' ');
# example 3
$n3 = stripos($s1, 'MON');
# example 4
$n4 = strripos($s1, 'MON');
# print
var_dump($n1, $n2, $n3, $n4);
