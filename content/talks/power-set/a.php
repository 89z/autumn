<?php
function powerset($a_in) {
   $a_out = [[]];
   foreach ($a_in as $a_elem) {
      foreach ($a_out as $a_comb) {
         $a_out[] = array_merge([$a_elem], $a_comb);
      }
   }
   return $a_out;
}
$a1 = [10, 11, 12];
$a2 = powerset($a1);
var_export($a2);
