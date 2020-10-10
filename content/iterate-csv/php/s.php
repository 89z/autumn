<?php

$s = <<<eof
Month,Day
January,Sunday
February,Monday
eof;

$s_row = strtok($s, "\n");
$a_head = str_getcsv($s_row);

while (true) {
   $s_row = strtok("\n");
   if ($s_row === false) {
      break;
   }
   $a_row = str_getcsv($s_row);
   $m_row = array_combine($a_head, $a_row);
   print_r($m_row);
}
