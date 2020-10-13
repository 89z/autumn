<?php
$o1 = date_create('2019-12-31');
$o2 = date_create();
$o = date_diff($o2, $o1);
$n = $o->days * 24 * 60 * 60 + $o->h * 60 * 60 + $o->i * 60 + $o->s;
var_dump($n);
