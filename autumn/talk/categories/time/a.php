<?php
$n = 1609459199;
# example 1
$date = new DateTime;
$date->setTimestamp($n);
$s1 = $date->format(DATE_W3C);
# example 2
$zone = new DateTimeZone('America/Chicago');
$date = new DateTime(timezone: $zone);
$date->setTimestamp($n);
$s2 = $date->format(DATE_W3C);
# print
var_dump(
   $s1 == '2020-12-31T23:59:59+00:00',
   $s2 == '2020-12-31T17:59:59-06:00'
);
