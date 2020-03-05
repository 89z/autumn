<?php
$n1 = 1000;
# example 1
$s1 = sprintf('%5d', $n1);
# example 2
$s2 = number_format($n1);
# print
var_dump($s1, $s2);
