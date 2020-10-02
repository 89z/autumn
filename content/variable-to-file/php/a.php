<?php
$s = "May\n";
# example 1
file_put_contents('a.txt', $s);
# example 2
$r = fopen('b.txt', 'w');
fwrite($r, $s);
# example 3
$r = fopen('c.txt', 'w');
fputs($r, $s);
