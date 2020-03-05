<?php
function u_group($s_in) {
   $n_neg = strrpos("-$s_in", '-');
   $n_dot = strpos("$s_in.", '.');
   $s_out = '';
   foreach (str_split($s_in) as $n_pos => $s_pos) {
      if ($n_neg < $n_pos && $n_pos < $n_dot && $n_pos % 3 == $n_dot % 3) {
         $s_out .= ',';
      }
      $s_out .= $s_pos;
   }
   return $s_out;
}
$a1 = ['123', '1234', '-123', '-1234', '123.4', '1234.5'];
foreach ($a1 as $s1) {
   var_dump(u_group($s1));
}
