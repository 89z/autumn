<?php
$s = '2020-12-31';
$d = date_create_from_format('!Y-m-d', $s);
var_dump($d);
