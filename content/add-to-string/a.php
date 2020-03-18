<?php
# example 1
$s1 = 'Sun';
$s1 = $s1 . 'day';
# example 2
$s2 = 'Sun';
$s2 .= 'day';
# example 3
$s3 = 'Sun';
$s3 = "${s3}day";
# example 4
$s4 = 'Sun';
$s4 = "$s4 Mon";
# print
var_dump($s1, $s2, $s3, $s4);
