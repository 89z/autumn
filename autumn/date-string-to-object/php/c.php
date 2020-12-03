<?php
$s = '2019-12-31';
$o = date_create_from_format('!Y-m-d', $s);
$n = date_timestamp_get($o);
var_dump($n == 1577750400);
