<?php
function f_start(...$a_arg) {
   $a_esc = array_map('escapeshellarg', $a_arg);
   $s_esc = implode(' ', $a_esc);
   return system($s_esc);
}
f_start('firefox', 'example.com');
