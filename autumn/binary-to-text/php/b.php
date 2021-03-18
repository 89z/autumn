<?php
$s = "\xa\xb\xc";
$t = bin2hex($s);
var_dump($t == '0a0b0c');
