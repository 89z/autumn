<?php
$s1 = 'Sunday';
# example 1
$s2 = substr($s1, 1, 1);
# example 2
$s3 = $s1[1];
# example 3
$s4 = substr($s1, 1);
# print
var_dump($s2 == 'u', $s3 == 'u', $s4 == 'unday');
