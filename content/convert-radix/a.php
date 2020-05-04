<?php
$s_dig = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
# example 1
function r64_encode($n_in) {
   global $s_dig;
   $s_out = '';
   for ($n_sh = $n_in; $n_sh > 0; $n_sh >>= 6) {
      $n_key = $n_sh & 63;
      $s_out = $s_dig[$n_key] . $s_out;
   }
   return $s_out;
}
# example 2
function r64_decode($s_in) {
   global $s_dig;
   $n_out = 0;
   for ($n_sh = 0, $n_key = -1; $n_sh < 36; $n_sh += 6, $n_key -= 1) {
      $s_val = $s_in[$n_key];
      $n_out |= strpos($s_dig, $s_val) << $n_sh;
   }
   return $n_out;
}
# print
$n1 = time();
$s1 = r64_encode($n1);
$n2 = r64_decode($s1);
var_dump($n1, $s1, $n2 == $n1);
