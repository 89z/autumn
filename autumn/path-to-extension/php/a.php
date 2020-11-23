<?php
$s = 'a.php';
# example 1
$s1 = pathinfo($s, PATHINFO_EXTENSION);
# example 2
$s2 = pathinfo($s)['extension'];
# print
var_dump($s1 == 'php', $s2 == 'php');
