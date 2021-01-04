<?php
$s = '2020-05-04';
$o = date_create_from_format('!Y-m-d', $s);
var_dump($o);
