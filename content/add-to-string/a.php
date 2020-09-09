<?php
$s = 'May';
# example 1
$s = $s . 'June';
# example 2
$s .= 'July';
# print
var_dump($s == 'MayJuneJuly');
