<?php
$o1 = date_create('2019-12-31');
$o2 = date_create();
$o3 = date_diff($o2, $o1);
$n1 = $o3->days;
var_dump($n1);
