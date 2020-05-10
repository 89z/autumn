<?php
$s1 = ' ab ';
# example 1
$s2 = chop($s1);
# example 2
$s3 = ltrim($s1);
# example 3
$s4 = rtrim($s1);
# example 4
$s5 = trim($s1);
# print
var_dump($s2 == ' ab', $s3 == 'ab ', $s4 == ' ab', $s5 == 'ab');
