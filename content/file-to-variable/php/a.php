<?php
# example 1
$a1 = file('index.md');
# example 2
$a2 = file('index.md', FILE_IGNORE_NEW_LINES);
# exmaple 3
$a3 = file('index.md', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
# print
var_dump($a1, $a2, $a3);
