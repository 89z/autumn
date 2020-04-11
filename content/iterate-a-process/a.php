<?php
$s1 = 'ag -V';
# example 1
$r1 = popen($s1, 'r');
while (true) {
   $s2 = fgets($r1);
   if (feof($r1)) {
      break;
   }
   var_dump($s2);
}
# example 2
exec($s1, $a1);
foreach ($a1 as $s2) {
   var_dump($s2);
}
