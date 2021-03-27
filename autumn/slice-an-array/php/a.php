<?php
$a = [0, 10, 20, 30, 40];

# example 1
$b = array_slice($a, 2, 2);
var_dump($b);

# example 2
$b = array_slice($a, 2);
var_dump($b);
