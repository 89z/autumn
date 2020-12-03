<?php
$o = date_create('2019-12-31');
$n = date_timestamp_get($o);
var_dump($n == 1577750400);
