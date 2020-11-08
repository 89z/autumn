<?php

function str_contains($s, $s2) {
   return strpos($s, $s2) !== false;
}

$b = str_contains('June', 'un');
var_dump($b);
