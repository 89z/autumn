<?php
$o = new DateTime('2020-05-04T03:02:01');
$n = $o->getTimestamp();
var_dump($n == 1588561321);
