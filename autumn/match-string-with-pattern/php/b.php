<?php

function pattern(string $pattern, string $subject): array|bool {
   $n = preg_match($pattern, $subject, $a);
   if ($n != 1) {
      return false;
   }
   return $a;
}

$a = pattern('/o./', 'south north');
var_dump($a);
