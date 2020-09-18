<?php
$nA = microtime(true);
echo "May\n";
$nB = microtime(true);
$nC = 3e6 - 1e6 * ($nB - $nA);
usleep($nC);
