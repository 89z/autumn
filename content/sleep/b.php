<?php
$n = microtime(true);
echo "May\n";
$n2 = microtime(true);
$n3 = 3e6 - 1e6 * ($n2 - $n);
usleep($n3);
