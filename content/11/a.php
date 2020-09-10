<?php
$r1 = fopen('a.txt', 'r');
while (true) {
   $s1 = fgets($r1);
   if (feof($r1)) {
      break;
   }
   var_dump($s1);
}
