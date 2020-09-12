<?php
$s = 'USERPROFILE';
# example 1
$s2 = getenv($s);
# example 2
$m = getenv();
$s3 = $m[$s];
# example 3
$s4 = $_SERVER[$s];
# print
var_dump($s2, $s3, $s4);
