<?php

function find_submatch(string $pat, string $sub): string {
   preg_match($pat, $sub, $a);
   if (count($a) < 2) {
      return '';
   }
   return $a[1];
}

$s = find_submatch('/a(..)/', 'January');
var_dump($s == 'nu');
