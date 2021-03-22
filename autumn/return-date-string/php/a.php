<?php
$z = new DateTimeZone('America/Chicago');
$t = new DateTime(timezone: $z);

# example 1
$s = $t->format(DATE_W3C);
var_dump($s);

# example 2
$s = $t->format('Y-m-d');
var_dump($s);
