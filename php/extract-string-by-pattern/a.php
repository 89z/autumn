<?php
$s = 'January';
# example 1
preg_match_all('/a./', $s, $a1);
# example 2
preg_match_all('/a(..)/', $s, $a2);
# print
var_dump($a1, $a2);
