<?php
date_default_timezone_set('America/Chicago');
$o = date_create();
date_timestamp_set($o, 1577858399);
# example 1
$s1 = date_format($o, DATE_W3C);
# example 2
$s2 = date_format($o, 'c');
# print
var_dump(
   $s1 == '2019-12-31T23:59:59-06:00', $s2 == '2019-12-31T23:59:59-06:00'
);
