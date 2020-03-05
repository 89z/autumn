<?php
# example 1
$a1 = scandir('.');
print_r($a1);
# example 2
$r1 = opendir('.');
while (true) {
   $s1 = readdir($r1);
   if ($s1 === false) {
      break;
   }
   var_dump($s1);
}
# example 3
$a2 = glob('a.*');
print_r($a2);
