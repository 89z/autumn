<?php
$o = new DateTime('2019-12-31');
$o2 = new DateTime;
$n = $o2->diff($o)->days;
var_dump($n);
