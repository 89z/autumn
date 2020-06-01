<?php
function f_escape_sh($s_in) {
   $s_out = '';
   $a_in = str_split($s_in);
   foreach ($a_in as $s_chr) {
      if ($s_chr == ' ') {
         $s_out .= '" "';
      } else if ($s_chr == '"') {
         $s_out .= '\^"';
      } else {
         $s_out .= '^' . $s_chr;
      }
   }
   return $s_out;
}
$a1 = ['less', '-V'];
$a2 = array_map('f_escape_sh', $a1);
$s1 = implode(' ', $a2);
system($s1);
