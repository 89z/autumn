<?php
$s1 = '/sunday/monday.txt';
# example 1
$s2 = dirname($s1);
# example 2
$s3 = basename($s1);
# example 3
$s4 = pathinfo($s1, PATHINFO_DIRNAME);
# example 4
$s5 = pathinfo($s1, PATHINFO_BASENAME);
# example 5
$s6 = pathinfo($s1, PATHINFO_EXTENSION);
# print
var_dump($s2, $s3, $s4, $s5, $s6);
