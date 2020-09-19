<?php
$s = '📗';
# example 1
$n1 = strlen($s);
# example 2
$n2 = iconv_strlen($s);
# example 3
$s3 = utf8_decode($s);
$n3 = strlen($s3);
# print
var_dump($n1 == 4, $n2 == 1, $n3 == 1);
