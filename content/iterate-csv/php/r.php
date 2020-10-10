<?php

function f($r_in): array {
   while (true) {
      $n_row = ftell($r_in);
      $a_row = fgetcsv($r_in);
      if (feof($r_in)) {
         return $a_out;
      }
      if ($n_row == 0) {
         $a_head = $a_row;
         continue;
      }
      $a_out[] = array_combine($a_head, $a_row);
   }
}

$r = fopen('a.csv', 'r');
$a = f($r);
print_r($a);
