<?php
date_default_timezone_set('America/Chicago');
$n1 = strtotime('2019-12-31');
# example 1
$s1 = date('Y-m-d', $n1);
# example 2
$s2 = strftime('%F', $n1);
# print
var_dump($s1, $s2);
