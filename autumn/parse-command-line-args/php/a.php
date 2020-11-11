<?php

if ($argc == 1) {
   echo 'add.php [flags] <stem>
-p string
   prefix
-s string
   suffix
';
   exit(1);
}

$m = getopt('p:s:', [], $n);
$pre_s = key_exists('p', $m) ? $m['p'] : '';
$suf_s = key_exists('s', $m) ? $m['s'] : '';
echo $pre_s . $argv[$n] . $suf_s, "\n";
