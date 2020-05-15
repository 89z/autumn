<?php
function f_escape_sh($s_all) {
   $a_enc = explode('"', $s_all);
   $a_tran = array_map(fn($s_one) => '"' . $s_one . '"', $a_enc);
   return implode('"', $a_tran);
}
