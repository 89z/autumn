<?php
$s = '11';
# example 1
$n1 = intval($s);
# example 2
$n2 = intval($s, 10);
# print
var_dump($n1 === 11, $n2 === 11);
