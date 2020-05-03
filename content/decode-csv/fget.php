<?php
$r_tab = fopen('a.csv', 'r');
$a_head = fgetcsv($r_tab);
while (true) {
   $a_row = fgetcsv($r_tab);
   if (feof($r_tab)) {
      break;
   }
   $m_row = array_combine($a_head, $a_row);
   echo $m_row['city'], "\n";
}
