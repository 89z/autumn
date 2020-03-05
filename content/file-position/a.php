<?php
$r1 = fopen('a.txt', 'r');
$s1 = fgets($r1);
# example 1
fseek($r1, 0);
$s2 = fgets($r1);
# example 2
rewind($r1);
$s3 = fgets($r1);
# print
var_dump($s2, $s3);
