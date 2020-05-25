<?php
$s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
# example 1
function r64_encode($n_in) {
   global $s_dig;
   $s_out = '';
   while ($n_in > 0) {
      $s_out = $s_dig[$n_in % 64] . $s_out;
      $n_in = intdiv($n_in, 64);
   }
   return $s_out;
}
# example 2
function r64_decode($s_in) {
   global $s_dig;
   $n_out = 0;
   $n_len = strlen($s_in);
   while ($n_len > 0) {
      $s_chr = $s_in[-$n_len];
      $n_out = $n_out * 64 + strpos($s_dig, $s_chr);
      $n_len--;
   }
   return $n_out;
}
# print
$n1 = time();
$s1 = r64_encode($n1);
$n2 = r64_decode($s1);
var_dump($n1, $s1, $n2 == $n1);
