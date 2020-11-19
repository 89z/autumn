<?php
$s = 'Sunday Monday';
# example 1
preg_match_all('/..n/', $s, $a);
var_dump($a[0][0] == 'Sun', $a[0][1] == 'Mon');
# example 2
preg_match_all('/(..)n/', $s, $a);
var_dump($a[1][0] == 'Su', $a[1][1] == 'Mo');
