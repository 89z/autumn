<?php
$a_tab = file('a.csv');
$s_head = $a_tab[0];
$a_head = str_getcsv($s_head);
$a_body = array_slice($a_tab, 1);

foreach ($a_body as $s_row) {
   $a_row = str_getcsv($s_row);
   $m_row = array_combine($a_head, $a_row);
   $s_city = $m_row['city'];
   var_dump($s_city);
}
