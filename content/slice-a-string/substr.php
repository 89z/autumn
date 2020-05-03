<?php
$s1 = 'Sunday';
# example 1
$s2 = substr($s1, 3);
# example 2
$s3 = substr($s1, 0, 3);
# example 3
$s4 = substr($s1, -3);
# print
var_dump($s2 == 'day', $s3 == 'Sun', $s4 == 'day');
