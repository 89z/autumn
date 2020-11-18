<?php

function format_number(float $n): string {
   $n2 = (int)(log10($n) / 3);
   return sprintf('%.3f', $n / 1e3 ** $n2) . ['', ' k', ' M', ' G'][$n2];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
