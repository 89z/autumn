<?php
$n = 9.9;
# exmaple 1
printf("%.0f\n", $n);
# example 2
fprintf(STDOUT, "%.0f\n", $n);
# example 3
$s = sprintf('%.0f', $n);
echo $s, "\n";
