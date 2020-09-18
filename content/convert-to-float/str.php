<?php
$s = '9.9';
# example 1
$n1 = +($s);
# example 2
$n2 = (float)($s);
# example 3
$n3 = floatval($s);
# print
var_dump($n1 === 9.9, $n2 === 9.9, $n3 === 9.9);
