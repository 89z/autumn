<?php
$s = '11';
# example 1
$n = +($s);
# example 2
$n2 = (int)($s);
# example 3
$n3 = intval($s);
# example 4
$n4 = intval($s, 10);
# print
var_dump($n === 11, $n2 === 11, $n3 === 11, $n4 === 11);
