<?php
function f_escape_sh($s_all) {
   $a_enc = explode('"', $s_all);
   $a_map = array_map(fn($s_one) => '"' . $s_one . '"', $a_enc);
   return implode('"', $a_map);
}
$a1 = ['less', '-V'];
$a2 = array_map('f_escape_sh', $a1);
$s1 = implode(' ', $a2);
system($s1);
