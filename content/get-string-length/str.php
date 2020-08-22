<?php
$s1 = '📕';
# example 1
$n1 = strlen($s1);
# example 2
$n2 = iconv_strlen($s1);
# example 3
$s2 = utf8_decode($s1);
$n3 = strlen($s2);
# print
var_dump($n1 == 4, $n2 == 1, $n3 == 1);
