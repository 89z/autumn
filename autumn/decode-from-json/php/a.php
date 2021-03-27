<?php
$s = '{"month": 12, "day": 31}';

# example 1
$o = json_decode($s);
$n = $o->day;
var_dump($n === 31);

# example 2
$m = json_decode($s, true);
$n = $m['day'];
var_dump($n === 31);
