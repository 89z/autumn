<?php
date_default_timezone_set('America/Chicago');
# example 1
$s1 = date('Y-m-d');
# example 2
$s2 = strftime('%F');
# print
var_dump($s1, $s2);
