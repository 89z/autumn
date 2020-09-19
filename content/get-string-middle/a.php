<?php
$s = 'June';
# example 1
$s1 = $s[1];
# example 2
$s2 = substr($s, 1, 1);
# example 3
$s3 = substr($s, 1, 2);
# print
var_dump($s1 == 'u', $s2 == 'u', $s3 == 'un');
