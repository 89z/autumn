<?php
$n1 = microtime(true);
echo "Sunday\n";
$n2 = microtime(true);
usleep(3 - $n2 - $n1);
