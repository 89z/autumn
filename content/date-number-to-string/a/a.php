<?php
date_default_timezone_set('America/Chicago');
$o = new DateTime;
$o->setTimestamp(1577858399);
# example 1
$s1 = $o->format(DATE_W3C);
# example 2
$s2 = $o->format('c');
# print
var_dump(
   $s1 == '2019-12-31T23:59:59-06:00', $s2 == '2019-12-31T23:59:59-06:00'
);
