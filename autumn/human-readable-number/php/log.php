<?php
declare(strict_types = 1);
error_reporting(E_ALL);

function f($n3) {
   $n2 = log($n3, 1024);
   $n = floor($n2);
   return sprintf('%.3f ', $n3 / 1024 ** $n) . ['', 'K', 'M', 'G'][$n];
}

$s = f(1264);
var_dump($s == '1.234 K');
