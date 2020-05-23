<?php
$m = ['Sunday' => 10];
if (! array_key_exists('Sunday', $m)) {
   echo "Error: None\n";
   exit(1);
}
echo $m['Sunday'], "\n";
