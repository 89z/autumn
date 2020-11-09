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
$s_pre = key_exists('p', $m) ? $m['p'] : '';
$s_suf = key_exists('s', $m) ? $m['s'] : '';
echo $s_pre . $argv[$n] . $s_suf, "\n";
