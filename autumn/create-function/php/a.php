<?php

# example 1
function f(int $d, int $e): int {
   return $d + $e;
}

# example 2
function g($d, $e) {
   return $d + $e;
}

# print
var_dump(f(4, 5) == 9);
var_dump(g(4, 5) == 9);
