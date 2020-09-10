<?php
$s = 'June';
# example 1
$s2 = $s[1];
# example 2
$s3 = substr($s, 1, 1);
# example 3
$s4 = substr($s, 1, 2);
# print
var_dump($s2 == 'u', $s3 == 'u', $s4 == 'un');
