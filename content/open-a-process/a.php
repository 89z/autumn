<?php
function u_open(...$a1) {
   $a2 = array_map('escapeshellarg', $a1);
   $s1 = implode(' ', $a2);
   return popen($s1, 'r');
}
$r1 = u_open('cal', '2019');
print_r($r1);
