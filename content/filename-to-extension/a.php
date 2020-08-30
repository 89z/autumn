<?php
$s = 'a.php';
# example 1
$s1 = pathinfo($s, PATHINFO_EXTENSION);
# example 2
$o = new SplFileInfo($s);
$s2 = $o->getExtension();
# print
var_dump($s1 == 'php', $s2 == 'php');
