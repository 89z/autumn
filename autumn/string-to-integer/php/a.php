<?php
$s = '100';
# example 1
$n1 = intval($s);
# example 2
$n2 = intval($s, 10);
# print
var_dump($n1 === 100, $n2 === 100);
