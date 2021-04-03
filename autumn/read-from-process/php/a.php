<?php
$p = popen('php -v', 'r');

while (true) {
   $s = fgets($p);
   if (feof($p)) {
      break;
   }
   var_dump($s);
}
