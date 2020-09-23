<?php
$n = 9;
# example 1
$s1 = (string)($n);
# example 2
$s2 = sprintf('%d', $n);
# example 3
$s3 = strval($n);
# print
var_dump($s1 === '9', $s2 === '9', $s3 === '9');
