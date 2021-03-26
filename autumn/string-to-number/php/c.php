<?php

# example 1
sscanf('100', '%d', $n);
var_dump($n === 100);

# example 2
sscanf('99.9', '%f', $n);
var_dump($n === 99.9);

# example 3
$n = sscanf('100', '%d')[0];
var_dump($n === 100);
