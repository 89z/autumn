<?php
$s = "\n\v\f\r";
$t = bin2hex($s);
var_dump($t == '0a0b0c0d');
