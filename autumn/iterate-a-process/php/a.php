<?php
declare(strict_types = 1);
$command = 'php -v';
$r = popen($command, 'r');

while (true) {
   $s = fgets($r);
   if (feof($r)) {
      break;
   }
   var_dump($s);
}
