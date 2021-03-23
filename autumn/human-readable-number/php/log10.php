<?php

function format_number(float $d): string {
   $e = (int)(log10($d) / 3);
   return sprintf('%.3f', $d / 1e3 ** $e) . ['', ' k', ' M', ' G'][$e];
}

$s = format_number(9012345678);
var_dump($s == '9.012 G');
