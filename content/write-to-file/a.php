<?php
$s1 = "Sunday\n";
# example 1
file_put_contents('a.txt', $s1);
# example 2
$r1 = fopen('b.txt', 'w');
fwrite($r1, $s1);
# example 3
$r2 = fopen('c.txt', 'w');
fputs($r2, $s1);
