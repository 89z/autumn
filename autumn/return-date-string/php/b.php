<?php
$z = timezone_open('America/Chicago');
$t = date_create(timezone: $z);

# example 1
$s = date_format($t, DATE_W3C);
var_dump($s);

# example 2
$s = date_format($t, 'Y-m-d');
var_dump($s);
