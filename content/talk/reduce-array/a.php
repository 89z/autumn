<?php
$a = ['May', 'June'];
$f = fn ($sa, $sc) => $sa . $sc;
$s = array_reduce($a, $f);
var_dump($s == 'MayJune');
