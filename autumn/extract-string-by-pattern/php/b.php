<?php

function find_submatch(string $pat, string $sub): string {
   $e = preg_match($pat, $sub, $a);
   if ($e != 1) {
      return '';
   }
   return $a[1];
}

$s = find_submatch('/a(..)/', 'January');
var_dump($s == 'nu');
