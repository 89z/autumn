<?php
$o = new DateTime('2019-12-31');
$n = $o->getTimestamp();
var_dump($n == 1577750400);
