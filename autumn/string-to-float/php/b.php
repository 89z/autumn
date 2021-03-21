<?php

# example 1
$n = sscanf('9.9', '%f')[0];
var_dump($n === 9.9);

# example 2
sscanf('9.9', '%f', $n);
var_dump($n === 9.9);
