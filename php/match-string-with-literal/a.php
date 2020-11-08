<?php

function str_starts_with(string $s, string $s2): bool {
   return strpos($s, $s2) === 0;
}

$b = str_starts_with('June', 'Ju');
var_dump($b);
