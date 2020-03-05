<?php
# example 1
$s1 = 'sun';
$s1 = $s1 . ' mon';
# example 2
$s2 = 'sun';
$s2 .= ' mon';
# example 3
$s3 = 'sun';
$s3 = "${s3} mon";
# example 4
$s4 = 'sun';
$s4 = "$s4 mon";
# print
var_dump($s1, $s2, $s3, $s4);
