<?php
$s = "\xa\xb\xc\xd";
$t = bin2hex($s);
var_dump($t == '0a0b0c0d');
