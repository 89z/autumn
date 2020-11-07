<?php
date_default_timezone_set('America/Chicago');
$o = date_create();
# example 1
$s1 = date_format($o, DATE_W3C);
# example 2
$s2 = date_format($o, 'Y-m-d');
# print
var_dump($s1, $s2);
