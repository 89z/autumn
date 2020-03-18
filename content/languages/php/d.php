<?php
$s1 = str_repeat('Sunday"', 2 ** 18);
$s1 .= file_get_contents('https://github.com/login');
$t1 = microtime(true);
$a1 = explode('"', $s1);
foreach ($a1 as $n1 => $s2) {
   if ($s2 == 'authenticity_token') {
      $s3 = $a1[$n1 + 2];
      break;
   }
}
$t2 = microtime(true);
echo $t2 - $t1, "\t", $s3, "\n";
