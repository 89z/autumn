<?php

if ($argc == 1) {
   echo 'add.php [flags] <stem>
-p string
   prefix
-s string
   suffix
';
   exit;
}

$m = getopt('p:s:', [], $n);
$pre = key_exists('p', $m) ? $m['p'] : '';
$suf = key_exists('s', $m) ? $m['s'] : '';
echo $pre . $argv[$n] . $suf, "\n";
