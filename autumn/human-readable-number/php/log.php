<?php

function format_number($n) {
   $n2 = log($n, 1024);
   $n3 = floor($n2);
   return sprintf('%.3f', $n / 1024 ** $n3) . ['', ' k', ' M', ' G'][$n3];
}

$s = format_number(1264);
var_dump($s == '1.234 k');
