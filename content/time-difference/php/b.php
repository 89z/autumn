<?php
$o = date_create('2019-12-31');
$o2 = date_create('2019-12-31T23:59:59');
$o3 = date_diff($o2, $o);
$s = date_interval_format($o3, '%H:%I:%S');
var_dump($s == '23:59:59');
