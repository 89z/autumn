<?php
$m = stat('a.php');
# example 1
$n1 = $m[9];
# example 2
$n2 = $m['mtime'];
# print
var_dump($n1, $n2);
