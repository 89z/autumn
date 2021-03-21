<?php
[$f, $i] = [7.5, 7];

# example 1
$n = $f / 2;
var_dump($n === 3.75);

# example 2
$n = (int)($f / 2);
var_dump($n === 3);

# example 3
$n = (int)($i / 2);
var_dump($n === 3);

# example 4
$n = $i / 2;
var_dump($n === 3.5);
