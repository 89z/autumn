<?php
$o1 = new DateTime('2019-12-31');
$o2 = new DateTime;
$o = $o2->diff($o1);
$n = $o->days * 24 * 60 * 60 + $o->h * 60 * 60 + $o->i * 60 + $o->s;
var_dump($n);
