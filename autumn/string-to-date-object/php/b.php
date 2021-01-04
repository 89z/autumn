<?php
$s = '2020-12-31';
$o = DateTime::createFromFormat('!Y-m-d', $s);
var_dump($o);
