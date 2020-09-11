<?php
$s = '9';
# example 1
$n = +($s);
# example 2
$n2 = (int)($s);
# example 3
$n3 = intval($s);
# example 4
$n4 = intval($s, 10);
# print
var_dump($n === 9, $n2 === 9, $n3 === 9, $n4 === 9);
