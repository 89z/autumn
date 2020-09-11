<?php
$s = '📗';
# example 1
$n = strlen($s);
# example 2
$n2 = iconv_strlen($s);
# example 3
$s2 = utf8_decode($s);
$n3 = strlen($s2);
# print
var_dump($n == 4, $n2 == 1, $n3 == 1);
