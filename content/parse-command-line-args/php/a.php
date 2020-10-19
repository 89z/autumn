<?php

if ($argc == 1) {
   echo <<<eof
slice.php [flags] <string>
-a int
   starting position
-b int
   substring length (default 1)
eof;
   exit(1);
}

$m = getopt('a:b:', [], $n);
$s_in = $argv[$n];
$n_start = array_key_exists('a', $m) ? +$m['a'] : 0;
$n_len = array_key_exists('b', $m) ? +$m['b'] : 1;
echo substr($s_in, $n_start, $n_len), "\n";
