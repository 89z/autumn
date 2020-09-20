<?php

class Shell {
   function escape(string $s_in): string {
      $s_brk = strpbrk($s_in, ' "&->^');
      if ($s_brk === false) {
         return $s_in;
      }
      return '"' . str_replace('"', '""', $s_in) . '"';
   }
}

$o = new Shell;
$s = $o->escape('google.com/search?tbm=vid&q=squarepusher');
system('waterfox ' . $s);
