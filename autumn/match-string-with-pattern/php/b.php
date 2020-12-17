<?php

function pattern(string $pattern, string $subject): array|bool {
   $n = preg_match($pattern, $subject, $a);
   if ($n != 1) {
      return false;
   }
   return $a;
}

$a = pattern('/..n/', 'Sunday Monday');
var_dump($a);
