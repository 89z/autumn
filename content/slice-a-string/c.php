<?php
$s1 = 'Sunday';
# example 1
$s2 = strstr($s1, 'd', true);
# example 2
$s3 = strstr($s1, 'd');
# print
var_dump($s2 == 'Sun', $s3 == 'day');
