<?php
$s = '2019-12-31';
$o = DateTime::createFromFormat('!Y-m-d', $s);
$n = $o->getTimestamp();
var_dump($n == 1577750400);
