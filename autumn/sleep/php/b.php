<?php
$n = microtime(true);
echo "March\n";
$n1 = microtime(true);
$n1a = 3e6 - 1e6 * ($n1 - $n);
usleep($n1a);
