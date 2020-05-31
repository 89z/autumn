<?php
function f_escape_sh($s_in) {
   # split on double quotes
   $a_in = explode('"', $s_in);
   # quote parts
   $a_out = array_map(fn($s_out) => '"' . $s_out . '"', $a_in);
   # join parts
   return implode('\^"', $a_out);
}
$a1 = ['less', '-V'];
$a2 = array_map('f_escape_sh', $a1);
$s1 = implode(' ', $a2);
system($s1);
