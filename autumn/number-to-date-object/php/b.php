<?php
$n = 0x5555_5555;
$d = new DateTime;
$d->setTimestamp($n);
var_dump($d);
