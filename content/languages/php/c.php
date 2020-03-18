<?php
$s1 = str_repeat('Sunday"', 2 ** 18);
$s1 .= file_get_contents('https://github.com/login');
$t1 = microtime(true);
$s2 = strtok($s1, '"');
while (true) {
   if ($s2 == 'authenticity_token') {
      strtok('"');
      $s3 = strtok('"');
      break;
   }
   $s2 = strtok('"');
}
$t2 = microtime(true);
echo $t2 - $t1, "\t", $s3, "\n";
