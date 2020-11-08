<?php

function str_contains(string $s, string $s2): bool {
   return strpos($s, $s2) !== false;
}

$b = str_contains('June', 'un');
var_dump($b);
