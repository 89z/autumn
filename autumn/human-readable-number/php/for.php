<?php

function format_number(float $n): string {
   for ($n2 = 0; $n >= 1e3; $n2++) {
      $n /= 1e3;
   }
   return sprintf('%.3f', $n) . ['', ' k', ' M', ' G'][$n2];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
