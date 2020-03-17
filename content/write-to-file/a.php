<?php
$s1 = "Sunday\n";
# example 1
file_put_contents('e1.txt', $s1);
# example 2
$r1 = fopen('e2.txt', 'w');
fwrite($r1, $s1);
# example 3
$r2 = fopen('e3.txt', 'w');
fputs($r2, $s1);
