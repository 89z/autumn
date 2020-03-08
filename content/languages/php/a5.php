<?php
$s1 = str_repeat('Sunday"', 2 ** 18);
$s1 .= file_get_contents('https://github.com/login');
$t1 = microtime(true);
$a1 = explode('"', $s1);
$n1 = array_search('authenticity_token', $a1);
$s2 = $a1[$n1 + 2];
$t2 = microtime(true);
echo $t2 - $t1, "\t", $s2, "\n";
