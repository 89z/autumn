<?php
$s = "\xa\xb\xc";
$t = unpack('H*', $s)[1];
var_dump($t == '0a0b0c');
