<?php
$s = 'January';
# example 1
$n1 = preg_match('/^J/', $s);
# example 2
$n2 = preg_match('/ja/i', $s);
# print
var_dump($n1 !== 0, $n2 !== 0);
