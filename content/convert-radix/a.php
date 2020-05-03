<?php
$s_safe = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
# example 1
function r64_encode($n_in) {
   global $s_safe;
   $s_out = '';
   for ($n_sh = $n_in; $n_sh > 0; $n_sh >>= 6) {
      $n_ind = $n_sh & 63;
      $s_out .= $s_safe[$n_ind];
   }
   return $s_out;
}
$s1 = r64_encode(1234567890);
# example 2
function r64_decode($s_in) {
   global $s_safe;
   $n_out = 0;
   for ($n_sh = 0, $n_in = 0; $n_sh < 36; $n_sh += 6, $n_in += 1) {
      $s_chr = $s_in[$n_in];
      $n_out |= strpos($s_safe, $s_chr) << $n_sh;
   }
   return $n_out;
}
$n1 = r64_decode('ibwB91');
# print
var_dump($s1 == 'ibwB91', $n1 == 1234567890);
