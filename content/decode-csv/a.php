<?php
$r1 = fopen('a.csv', 'r');
$a1 = [];
while (true) {
   $a2 = fgetcsv($r1);
   if (feof($r1)) {
      break;
   }
   if ($a1 == []) {
      $a1 = $a2;
      continue;
   }
   $m1 = array_combine($a1, $a2);
   echo $m1['city'], "\n";
}
