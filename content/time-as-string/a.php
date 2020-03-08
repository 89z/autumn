<?php
# example 1
$s1 = strftime('%F');
# example 2
$s2 = date(DATE_W3C, 86400);
# example 3
$s3 = date('c');
# example 4
date_default_timezone_set(getenv('TZ'));
$s4 = date('c');
# print
var_dump($s1, $s2, $s3, $s4);
