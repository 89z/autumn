<?php

function f_escape_sh(string $s_in): string {
   $s_brk = strpbrk($s_in, ' "&->^');
   if ($s_brk === false) {
      return $s_in;
   }
   return '"' . str_replace('"', '""', $s_in) . '"';
}

$s = f_escape_sh('a b.txt');
system('less ' . $s);
