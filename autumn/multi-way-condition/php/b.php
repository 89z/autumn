<?php
$s = "\x43";

switch ($s) {
case 'A':
   $n = 0x41;
   break;
case 'B':
case 'b':
   $n = 0x42;
   break;
default:
   $n = ord($s);
}

var_dump($n == 0x43);
