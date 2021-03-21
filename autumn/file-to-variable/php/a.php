<?php
# example 1
$a = file('a.php');
# example 2
$b = file('a.php', FILE_IGNORE_NEW_LINES);
# exmaple 3
$c = file('a.php', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
# print
var_dump($a, $b, $c);
