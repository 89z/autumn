<?php
$a = ['Sun', 'Mon', 'Tue'];
# example 1
$s = $a[0];
# example 2
$s2 = reset($a);
# example 3
$s3 = end($a);
# example 4
$a2 = array_slice($a, 0, 2);
# example 5
$a3 = array_slice($a, -2);
# print
var_dump($s, $s2, $s3, $a2, $a3);
