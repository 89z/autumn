<?php
date_default_timezone_set('America/Chicago');
# example 1
$s1 = strftime('%F');
# example 2
$s2 = strftime('%FT%T');
# print
var_dump($s1, $s2);
