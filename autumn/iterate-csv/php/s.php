<?php

$s = <<<eof
Month,Day
January,Sunday
February,Monday
eof;

$row_s = strtok($s, "\n");
$head_a = str_getcsv($row_s);

while (true) {
   $row_s = strtok("\n");
   if ($row_s === false) {
      break;
   }
   $row_a = str_getcsv($row_s);
   $row_m = array_combine($head_a, $row_a);
   print_r($row_m);
}
