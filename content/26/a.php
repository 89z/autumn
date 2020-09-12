<?php
$n = 1000;
# example 1
$s = (string)($n);
# example 2
$s2 = sprintf('%d', $n);
# example 3
$s3 = strval($n);
# example 4
$s4 = number_format($n);
# print
var_dump($s === '1000', $s2 === '1000', $s3 === '1000', $s4 === '1,000');
