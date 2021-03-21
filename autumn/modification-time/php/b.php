<?php
$m = stat('a.php');

# example 1
$n = $m[9];
var_dump($n);

# example 2
$n = $m['mtime'];
var_dump($n);
