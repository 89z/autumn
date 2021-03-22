<?php
[$f, $i] = ['7.5', '7'];

# example 1
$s = bcdiv($f, '2');
var_dump($s === '3');

# exmaple 2
$s = bcdiv($i, '2');
var_dump($s === '3');
