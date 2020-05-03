<?php
# begin
function r64_encode($n_in) {
   $a_safe = array_merge(
      range(0, 9),
      range('a', 'z'),
      range('A', 'Z'),
      ['-', '_']
   );
   $s_out = '';
   while ($n_in > 0) {
      $s_out .= $a_safe[$n_in & 63];
      $n_in >>= 6;
   }
   return $s_out;
}
# end
$s1 = r64_encode(1234567890);
var_dump($s1);
