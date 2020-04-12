<?php
# example 1
$r_file = fopen('a.csv', 'r');
$m_head = [];
while (true) {
   $a_body = fgetcsv($r_file);
   if (feof($r_file)) {
      break;
   }
   if (count($m_head) == 0) {
      $m_head = array_flip($a_body);
      continue;
   }
   echo $a_body[$m_head['city']], "\n";
}
# exmaple 2
$r_file = fopen('a.csv', 'r');
$a_head = [];
while (true) {
   $a_body = fgetcsv($r_file);
   if (feof($r_file)) {
      break;
   }
   if (count($a_head) == 0) {
      $a_head = $a_body;
      continue;
   }
   $m_body = array_combine($a_head, $a_body);
   echo $m_body['city'], "\n";
}
