<?php
# example 1
$o1 = new DateTime;
# exmaple 2
$o2 = new DateTime();
# example 3
$o3 = new DateTime('now');
# exmaple 4
$o = new DateTimeZone('America/Chicago');
$o4 = new DateTime(timezone: $o);
# print
var_dump($o1, $o2, $o3, $o4);
