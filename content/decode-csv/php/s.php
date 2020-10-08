<?php

function f(string $s_in): array {
   $a_in = explode("\n", $s_in);
   foreach ($a_in as $n_row => $s_row) {
      $a_row = str_getcsv($s_row);
      if ($n_row == 0) {
         $a_head = $a_row;
         continue;
      }
      $a_out[] = array_combine($a_head, $a_row);
   }
   return $a_out;
}

$s = <<<eof
Month,Day
January,Sunday
February,Monday
eof;

$a = f($s);
print_r($a);
