<?php
$n = 1577858399;
# example 1
$s1 = date(DATE_W3C, $n);
# example 2
date_default_timezone_set('America/Chicago');
$s2 = date(DATE_W3C, $n);
# print
var_dump(
   $s1 == '2020-01-01T05:59:59+00:00',
   $s2 == '2019-12-31T23:59:59-06:00'
);
