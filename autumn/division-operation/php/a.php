<?php
[$f, $i] = ['7.5', '7'];
# example 1
$s1 = bcdiv($f, '2');
# exmaple 2
$s2 = bcdiv($i, '2');
# print
var_dump($s1 === '3', $s2 === '3');
