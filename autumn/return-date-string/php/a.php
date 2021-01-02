<?php
$zone = new DateTimeZone('America/Chicago');
$date = new DateTime(timezone: $zone);
# example 1
$s1 = $date->format(DATE_W3C);
# example 2
$s2 = $date->format('Y-m-d');
# print
var_dump($s1, $s2);
