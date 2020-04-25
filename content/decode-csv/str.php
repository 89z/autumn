<?php
$a_table = file('a.csv');
$s_head = $a_table[0];
$a_head = str_getcsv($s_head);
$a_body = array_slice($a_table, 1);
foreach ($a_body as $s_row) {
   $a_row = str_getcsv($s_row);
   $m_row = array_combine($a_head, $a_row);
   echo $m_row['city'], "\n";
}
