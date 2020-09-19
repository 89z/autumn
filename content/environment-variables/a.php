<?php
$s = 'USERPROFILE';
# example 1
$s1 = getenv($s);
# example 2
$m = getenv();
$s2 = $m[$s];
# example 3
$s3 = $_SERVER[$s];
# print
$s = 'C:\\Users\\Steven';
var_dump($s1 == $s, $s2 == $s, $s3 == $s);
