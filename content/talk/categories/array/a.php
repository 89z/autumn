<?php
# example 1
$a1 = [10, 11];
$a2 = [10, 11];
var_dump($a1 == $a2);
# example 2
$a1 = [true, false];
$a2 = [1, 0];
var_dump($a1 == $a2);
# example 3
$a1 = [10, 11, 12, 13];
$a2 = [[10, 11], [12, 13]];
var_dump($a1 != $a2);
