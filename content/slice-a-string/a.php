<?php
$s1 = 'Sunday';
# example 1
$s2 = $s1[0];
# example 2
$s3 = substr($s1, 0, 1);
# example 3
$s4 = substr($s1, 3);
# example 4
$s5 = substr($s1, -3);
# example 5
$s6 = $s1[-1];
# print
var_dump($s2 == 'S', $s3 == 'Sun', $s4 == 'day', $s5 == 'day', $s6 == 'y');
