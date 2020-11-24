<?php
$r = fopen('a.csv', 'r');
$head_a = fgetcsv($r);

while (true) {
   $row_a = fgetcsv($r);
   if (feof($r)) {
      break;
   }
   $row_m = array_combine($head_a, $row_a);
   print_r($row_m);
}
