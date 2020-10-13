<?php
$o = date_create('2019-12-31');
$o2 = date_create();
$n = date_diff($o2, $o)->days;
var_dump($n);
