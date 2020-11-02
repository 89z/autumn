<?php
$n = 1 + 2;

switch ($n) {
case 5:
   $s = 'five';
   break;
case 4:
case 3:
   $s = 'four or three';
   break;
default:
   $s = strval($n);
}

var_dump($s == 'four or three');
