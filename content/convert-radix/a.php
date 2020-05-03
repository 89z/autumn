<?php
# head
function radix64($n_in) {
   $s_safe = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
   $s_out = '';
   while ($n_in > 0) {
      $s_out .= $s_safe[$n_in & 63];
      $n_in >>= 6;
   }
   return $s_out;
}
# body
$s1 = radix64(1588337932);
var_dump($s1);
