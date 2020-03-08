<?php
$s1 = file_get_contents('https://github.com/login');
$s2 = strtok($s1, '"');
while (true) {
   if ($s2 == 'authenticity_token') {
      strtok('"');
      $s3 = strtok('"');
      break;
   }
   $s2 = strtok('"');
}
echo $s3, "\n";
