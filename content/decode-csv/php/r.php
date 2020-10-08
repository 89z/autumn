<?php
$r_in = fopen('a.csv', 'r');

while (true) {
   $n_row = ftell($r_in);
   $a_row = fgetcsv($r_in);
   if (feof($r_in)) {
      break;
   }
   if ($n_row == 0) {
      $a_head = $a_row;
      continue;
   }
   $a_out[] = array_combine($a_head, $a_row);
}

print_r($a_out);
