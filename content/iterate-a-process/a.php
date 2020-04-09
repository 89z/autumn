<?php
$r1 = popen('locale', 'r');
while (true) {
   $s1 = fgets($r1);
   if (feof($r1)) {
      break;
   }
   echo $s1;
}
