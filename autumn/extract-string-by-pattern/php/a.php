<?php

function find_string(string $pat, string $sub): string {
   $e = preg_match($pat, $sub, $a);
   if ($e != 1) {
      return '';
   }
   return $a[0];
}

$s = find_string('/a../', 'January');
var_dump($s == 'anu');
