<?php

function format_number($n) {
   for ($n2 = 0; $n > 1024; $n2++) {
      $n /= 1024;
   }
   return sprintf('%.3f', $n) . ['', ' k', ' M', ' G'][$n2];
}

$s = format_number(1264);
var_dump($s == '1.234 k');
