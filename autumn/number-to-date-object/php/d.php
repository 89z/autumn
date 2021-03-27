<?php
$n = 1609480799;
$d = date_create();
date_timestamp_set($d, $n);
var_dump($d);
