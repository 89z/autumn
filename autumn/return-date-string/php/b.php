<?php
$zone = timezone_open('America/Chicago');
$date = date_create(timezone: $zone);
# example 1
$s1 = date_format($date, DATE_W3C);
# example 2
$s2 = date_format($date, 'Y-m-d');
# print
var_dump($s1, $s2);
