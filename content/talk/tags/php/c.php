<?php
$s1 = getenv('PUBLIC');
$o1 = dir($s1);
$s2 = $o1->path;
var_dump($s2); # "C:\Users\Public"
