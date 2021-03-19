<?php
$s = "\n\v\f\r";
$t = unpack('H*', $s)[1];
var_dump($t == '0a0b0c0d');
