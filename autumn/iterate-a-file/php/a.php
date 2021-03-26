<?php
$f = fopen('a.php', 'r');

while (true) {
   $s = fgets($f);
   if (feof($f)) {
      break;
   }
   var_dump($s);
}
