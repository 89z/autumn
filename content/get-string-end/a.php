<?php
$s = 'Sunday';
# example 1
$s1 = substr($s, 1);
# example 2
$s2 = substr($s, -1);
# example 3
$s3 = $s[-1];
# print
var_dump($s1 == 'unday', $s2 == 'y', $s3 == 'y');
