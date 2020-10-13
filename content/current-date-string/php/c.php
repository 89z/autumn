<?php
date_default_timezone_set('America/Chicago');
# example 1
$s1 = date(DATE_W3C);
# example 2
$s2 = date('Y-m-d');
# print
var_dump($s1, $s2);
