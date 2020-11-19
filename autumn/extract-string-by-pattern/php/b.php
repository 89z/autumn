<?php
$s = 'Sunday Monday';
# example 1
preg_match('/..n/', $s, $a);
var_dump($a[0] == 'Sun');
# example 2
preg_match('/(..)n/', $s, $a);
var_dump($a[1] == 'Su');
