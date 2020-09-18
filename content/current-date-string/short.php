<?php
date_default_timezone_set('America/Chicago');
# example 1
$s1 = date('Y-m-d');
# example 2
$o = new DateTime;
$s2 = $o->format('Y-m-d');
# example 3
$s3 = strftime('%F');
# print
var_dump($s1, $s2, $s3);
