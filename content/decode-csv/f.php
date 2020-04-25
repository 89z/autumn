<?php
$r_table = fopen('a.csv', 'r');
$a_head = fgetcsv($r_table);
while (true) {
   $a_row = fgetcsv($r_table);
   if (feof($r_table)) {
      break;
   }
   $m_row = array_combine($a_head, $a_row);
   echo $m_row['city'], "\n";
}
