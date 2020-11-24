<?php
# example 1
$n1 = time();
touch('a.php', $n1);
# example 2
$n2 = filemtime('a.php');
# print
var_dump($n1 == $n2);
