<?php
$a = ['May', 'June'];
$f = fn ($sa, $sc) => $sa . $sc;
$s = '';

foreach ($a as $sc) {
   $s = $f($s, $sc);
}

var_dump($s == 'MayJune');
