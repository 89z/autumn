<?php
# example 1
$b1 = strpos('Sunday', 'und') !== false;
# example 2
$b2 = strpos('Sunday', 'Sun') === 0;
# example 3
$s1 = explode('day', 'Sunday');
$b3 = end($a1) === '';
# example 4
$b4 = 'Sunday'[0] == 'S';
# print
var_dump($b1, $b2, $b3, $b4);
