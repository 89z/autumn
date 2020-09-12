<?php
$s = '2.9';
# example 1
$n = +($s);
# example 2
$n2 = (float)($s);
# example 3
$n3 = floatval($s);
# print
var_dump($n === 2.9, $n2 === 2.9, $n3 === 2.9);
