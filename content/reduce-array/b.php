<?php
function f(string $sa, string $sc): string {
   return $sa . $sc;
}

$a = ['May', 'June'];
$s = '';

foreach ($a as $sc) {
   $s = f($s, $sc);
}

var_dump($s == 'MayJune');
