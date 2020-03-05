<?php
$s1 = 'sunday';
# example 1
$s2 = $s1[3];
# example 2
$s3 = $s1[-3];
# example 3
$s4 = substr($s1, 3);
# example 4
$s5 = substr($s1, 0, 3);
# example 5
$s6 = substr($s1, -3);
# print
var_dump($s2, $s3, $s4, $s5, $s6);
