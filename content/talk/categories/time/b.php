<?php
$n = 365 * 24 * 60 * 60;
$o = new DateTime;
$o->setTimestamp($n);
$s = $o->format('Y-m-d');
var_dump($s);
