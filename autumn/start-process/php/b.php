<?php

function shell_escape(string $in): string {
   $brk = strpbrk($in, ' "&->^');
   if ($brk === false) {
      return $in;
   }
   return '"' . str_replace('"', '""', $in) . '"';
}

$s = shell_escape('google.com/search?tbm=vid&q=squarepusher');
system('waterfox ' . $s);
