<?php
$n = 0x5555_5555;
$d = date_create();
date_timestamp_set($d, $n);
var_dump($d);
