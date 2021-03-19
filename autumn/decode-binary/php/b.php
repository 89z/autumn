<?php
$s = '0a0b0c0d';
$t = hex2bin($s);
var_dump($t == "\xa\xb\xc\xd");
