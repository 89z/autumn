<?php

function shell_escape(string $s_in): string {
   $s_brk = strpbrk($s_in, ' "&->^');
   if ($s_brk === false) {
      return $s_in;
   }
   return '"' . str_replace('"', '""', $s_in) . '"';
}

$s = shell_escape('google.com/search?tbm=vid&q=squarepusher');
system('waterfox ' . $s);
