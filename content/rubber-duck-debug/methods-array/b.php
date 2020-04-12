<?php
foreach ($a1 as $n1 => $a2) {
   if ($n1 == 0) {
      $a3 = $a2;
   } else {
      $n2 = array_search('fri', $a3);
      echo $a2[$n2], "\n";
   }
}
