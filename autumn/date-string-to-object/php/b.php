<?php
$s = '2019-12-31';
$o = DateTime::createFromFormat('!Y-m-d', $s);
var_dump($o);
