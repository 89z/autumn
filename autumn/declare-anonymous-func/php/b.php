<?php

# example 1
$f = fn (int $d, int $e): int => $d + $e;
var_dump($f(4, 5) == 9);

# example 2
$f = fn ($d, $e) => $d + $e;
var_dump($f(4, 5) == 9);
