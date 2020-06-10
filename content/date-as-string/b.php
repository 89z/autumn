<?php
date_default_timezone_set('America/Chicago');
$n = strtotime('2019-12-31');
# example 1
$s1 = date('Y-m-d', $n);
# example 2
$s2 = strftime('%F', $n);
# print
var_dump($s1, $s2);
