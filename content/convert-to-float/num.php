<?php
$n = 10;
# example 1
$n1 = (float)($n);
# example 2
$n2 = floatval($n);
# print
var_dump($n1 === 10.0, $n2 === 10.0);
