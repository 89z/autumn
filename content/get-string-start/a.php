<?php
$s = 'May';
# example 1
$s1 = substr($s, 0, 1);
# example 2
$s2 = $s[0];
# example 3
$s3 = substr($s, 0, 2);
# print
var_dump($s1 == 'M', $s2 == 'M', $s3 == 'Ma');
