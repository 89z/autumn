<?php
$n = 1609480799;
$o = date_create();
date_timestamp_set($o, $n);
var_dump($o);
