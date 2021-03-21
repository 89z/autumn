<?php

# example 1
$t = new DateTime;
var_dump($t);

# exmaple 2
$t = new DateTime();
var_dump($t);

# example 3
$t = new DateTime('now');
var_dump($t);

# exmaple 4
$z = new DateTimeZone('America/Chicago');
$t = new DateTime(timezone: $z);
var_dump($t);
