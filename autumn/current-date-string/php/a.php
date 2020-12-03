<?php
$zone_o = new DateTimeZone('America/Chicago');
$date_o = new DateTime(timezone: $zone_o);
# example 1
$s1 = $date_o->format(DATE_W3C);
# example 2
$s2 = $date_o->format('Y-m-d');
# print
var_dump($s1, $s2);
