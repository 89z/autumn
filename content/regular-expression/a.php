<?php
$s1 = 'Sunday Monday';
# example 1
preg_match('/Mon.*/', $s1, $a1);
$s2 = $a1[0];
# example 2
preg_match('/Mon(.*)/', $s1, $a2);
$s3 = $a2[1];
# print
var_dump($s2, $s3);
