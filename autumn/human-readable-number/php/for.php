<?php

function format_number(float $d): string {
   $e = 0;
   for ($f = $d; $f >= 1e3; $f /= 1e3) {
      $e++;
   }
   return sprintf('%.3f', $f) . ['', ' k', ' M', ' G'][$e];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
