<?php

if ($argc == 1) {
   echo 'cat.php [flags] <input>
-s string
   start
-e string
   end
';
   exit(1);
}

$m = getopt('s:e:', [], $n);
$s_start = key_exists('s', $m) ? $m['s'] : '';
$s_end = key_exists('e', $m) ? $m['e'] : '';
echo $s_start, $argv[$n], $s_end, "\n";
