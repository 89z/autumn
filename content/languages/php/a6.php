<?php
$s1 = str_repeat('Sunday ', 2 ** 18);
$s1 .= file_get_contents('https://github.com/login');
$t1 = microtime(true);
preg_match('/name="authenticity_token" value="([^"]*)"/', $s1, $a1);
$s2 = $a1[1];
$t2 = microtime(true);
echo $t2 - $t1, "\t", $s2, "\n";
