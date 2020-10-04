<?php
$s = 'php -v';
# example 1
$r = popen($s, 'r');
while (true) {
   $s1 = fgets($r);
   if (feof($r)) {
      break;
   }
   var_dump($s1);
}
# example 2
exec($s, $a);
foreach ($a as $s2) {
   var_dump($s2);
}
