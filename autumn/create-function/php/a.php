<?php

# example 1
$f = function (int $d, int $e): int {
   return $d + $e;
};

var_dump($f(4, 5) == 9);

# example 2
$f = function ($d, $e) {
   return $d + $e;
};

var_dump($f(4, 5) == 9);
