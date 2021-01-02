<?php
$s = '2020-12-31';
$o = date_create_from_format('!Y-m-d', $s);
var_dump($o);
