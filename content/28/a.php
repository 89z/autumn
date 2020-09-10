<?php
date_default_timezone_set('America/Chicago');
# example 1
$s1 = date('Y-m-d');
# example 2
$s2 = strftime('%F');
# example 3
$o = date_create();
$s3 = $o->format('Y-m-d');
# print
var_dump($s1, $s2, $s3);
