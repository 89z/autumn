<?php
$s = '2020-05-04';
$o = DateTime::createFromFormat('!Y-m-d', $s);
var_dump($o);
