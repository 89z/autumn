<?php
$zone_o = timezone_open('America/Chicago');
$date_o = date_create(timezone: $zone_o);
# example 1
$s1 = date_format($date_o, DATE_W3C);
# example 2
$s2 = date_format($date_o, 'Y-m-d');
# print
var_dump($s1, $s2);
