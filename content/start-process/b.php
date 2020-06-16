<?php
function f_escape_sh($s_in) {
   $s_out = preg_replace('/"/', '""', $s_in);
   if ($s_out != $s_in) {
      return '"' . $s_out . '"';
   }
   $n_mat = preg_match('/[ &>^-]/', $s_in);
   if ($n_mat == 1) {
      return '"' . $s_out . '"';
   }
   return $s_out;
}
$s = f_escape_sh('a b.txt');
system('less ' . $s);
