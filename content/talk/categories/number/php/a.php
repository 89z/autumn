<?php
declare(strict_types = 1);
error_reporting(E_ALL);
$nX = 7;
$nY = 19;
# example 1
$n1 = $nX ** $nY;
# exmaple 2
$n2 = pow($nX, $nY);
# print
var_dump($n1, $n2);
