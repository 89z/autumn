<?php
# begin
function r64_decode($s_in) {
   $s_safe = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
   $n_out = 0;
   for ($n_sh = 0, $n_in = 0; $n_sh < 36; $n_sh += 6, $n_in += 1) {
      $s_chr = $s_in[$n_in];
      $n_out |= strpos($s_safe, $s_chr) << $n_sh;
   }
   return $n_out;
}
# end
$n1 = r64_decode('ibwB91');
var_dump($n1);
