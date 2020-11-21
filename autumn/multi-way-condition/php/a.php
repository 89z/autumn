<?php
$n = 0x63 - 0x20;

switch ($n) {
case 0x41:
   $s = 'A';
case 0x42:
case 0x62:
   $s = 'B';
default:
   $s = chr($n);
}

var_dump($s == 'C');
