<?php

function str_starts_with($s, $s2) {
   return strpos($s, $s2) === 0;
}

$b = str_starts_with('June', 'Ju');
var_dump($b);
