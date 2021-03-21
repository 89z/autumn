<?php
$m = stat('a.php');

# example 1
$n = $m[7];
var_dump($n);

# example 2
$n = $m['size'];
var_dump($n);
