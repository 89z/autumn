<?php
$s = '{"month": 12, "day": 31}';
# example 1
$o = json_decode($s);
# example 2
$m = json_decode($s, true);
# print
var_dump($o->day === 31, $m['day'] === 31);
