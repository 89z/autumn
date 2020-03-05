<?php
# example 1
$b1 = strpos('sunday', 'und') !== false;
# example 2
$b2 = strpos('sunday', 'sun') === 0;
# example 3
$a1 = explode('day', 'sunday');
$b3 = end($a1) === '';
# example 4
$b4 = 'sunday'[0] == 's';
# print
var_dump($b1, $b2, $b3, $b4);
