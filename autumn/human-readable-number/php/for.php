<?php

function format_number(float $n): string {
   $n2 = 0;
   for ($n3 = $n; $n3 >= 1e3; $n3 /= 1e3) {
      $n2++;
   }
   return sprintf('%.3f', $n3) . ['', ' k', ' M', ' G'][$n2];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
