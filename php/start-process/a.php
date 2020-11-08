<?php

# example 1
echo "BEGIN\n";
system('pipe');

# example 2
echo "BEGIN\n";
$a = [];
$r = proc_open('pipe', $a, $a);
proc_close($r);
