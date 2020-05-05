<?php
$s1 = 'Sunday';
# example 1
$s2 = substr($s1, 0, 1);
# example 2
$s3 = $s1[0];
# example 3
$s4 = substr($s1, 1, 2);
# example 4
$s5 = substr($s1, -1);
# example 5
$s6 = $s1[-1];
# print
var_dump($s2 == 'S', $s3 == 'S', $s4 == 'un', $s5 == 'y', $s6 =='y');
