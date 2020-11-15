<?php

function f($n) {
   for ($n2 = 0; $n > 1000; $n2++) {
      $n /= 1000;
   }
   return sprintf('%.3f ', $n) . ['', 'K', 'M', 'G'][$n2];
}

$s = f(1234);
var_dump($s == '1.234 K');
