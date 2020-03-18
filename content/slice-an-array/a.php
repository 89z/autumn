<?php
$a1 = ['Sun', 'Mon', 'Tue'];
# example 1
$s1 = $a1[0];
# example 2
$s2 = reset($a1);
# example 3
$s3 = end($a1);
# example 4
$a2 = array_slice($a1, 0, 2);
# example 5
$a3 = array_slice($a1, -2);
# print
var_dump($s1, $s2, $s3, $a2, $a3);
