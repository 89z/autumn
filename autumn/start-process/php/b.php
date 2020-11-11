<?php

function shell_escape(string $in_s): string {
   $brk_s = strpbrk($in_s, ' "&->^');
   if ($brk_s === false) {
      return $in_s;
   }
   return '"' . str_replace('"', '""', $in_s) . '"';
}

$s = shell_escape('google.com/search?tbm=vid&q=squarepusher');
system('waterfox ' . $s);
