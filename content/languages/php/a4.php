<?php
$s1 = file_get_contents('https://github.com/login');
$a1 = explode('"', $s1);
foreach ($a1 as $n1 => $s2) {
   if ($s2 == 'authenticity_token') {
      $s3 = $a1[$n1 + 2];
      break;
   }
}
echo $s3, "\n";
