<?php
# head
function radix64($n_in) {
   $a_safe = array_merge(
      range(0, 9),
      range('a', 'z'),
      range('A', 'Z'),
      ['-', '_']
   );
   $s_safe = implode('', $a_safe);
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
