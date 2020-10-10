<?php
$r = fopen('a.csv', 'r');
$a_head = fgetcsv($r);

while (true) {
   $a_row = fgetcsv($r);
   if (feof($r)) {
      break;
   }
   $m_row = array_combine($a_head, $a_row);
   print_r($m_row);
}
