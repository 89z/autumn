<?php
date_default_timezone_set('America/Chicago');
$o = new DateTime;
# example 1
$s1 = $o->format(DATE_W3C);
# example 2
$s2 = $o->format('Y-m-d');
# print
var_dump($s1, $s2);
