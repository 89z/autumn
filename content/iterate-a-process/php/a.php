<?php
$s = 'php -v';
$r = popen($s, 'r');

while (true) {
   $s1 = fgets($r);
   if (feof($r)) {
      break;
   }
   var_dump($s1);
}
