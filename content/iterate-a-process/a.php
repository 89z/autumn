<?php
$p1 = popen('cal', 'r');
while (true) {
   $s1 = fgets($p1);
   if (feof($p1)) {
      break;
   }
   echo $s1;
}
