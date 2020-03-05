<?php
foreach ($a1 as $n1 => $a2) {
   if ($n1 == 0) {
      $a3 = $a2;
   } else {
      $a4 = array_combine($a3, $a2);
      echo $a4['fri'], "\n";
   }
}
