<?php
$s = 'May';
# example 1
$s2 = substr($s, 0, 1);
# example 2
$s3 = $s[0];
# example 3
$s4 = substr($s, 0, 2);
# print
var_dump($s2 == 'M', $s3 == 'M', $s4 == 'Ma');
