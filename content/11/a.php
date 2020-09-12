<?php
$r = fopen('a.txt', 'r');

while (true) {
   $s = fgets($r);
   if (feof($r)) {
      break;
   }
   var_dump($s);
}
