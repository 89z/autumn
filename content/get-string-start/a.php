<?php
$s1 = 'Sunday';
# example 1
$s2 = substr($s1, 0, 1);
# example 2
$s3 = $s1[0];
# example 3
$s4 = substr($s1, 0, 2);
# print
var_dump($s2 == 'S', $s3 == 'S', $s4 == 'Su');
