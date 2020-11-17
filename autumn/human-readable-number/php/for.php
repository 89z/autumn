<?php

function format_number($n) {
   for ($n2 = 0; $n > 1000; $n2++) {
      $n /= 1000;
   }
   return sprintf('%.3f', $n) . ['', ' k', ' M', ' G'][$n2];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
