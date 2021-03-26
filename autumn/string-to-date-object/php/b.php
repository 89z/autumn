<?php
$s = '2020-12-31';
$d = DateTime::createFromFormat('!Y-m-d', $s);
var_dump($d);
