<?php
$s = 'Sun';
switch ($s) {
case 'Fri':
   $n = 1;
   break;
case 'Sat':
case 'Sun':
   $n = 2;
   break;
default:
   $n = 0;
}
var_dump($n);
