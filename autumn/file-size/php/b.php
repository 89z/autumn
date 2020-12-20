<?php
$m = stat('a.php');
# example 1
$n1 = $m[7];
# example 2
$n2 = $m['size'];
# print
var_dump($n1, $n2);
