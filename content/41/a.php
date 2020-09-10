<?php
$a = ['May', 'June'];
# example 1
$s = implode(',', $a);
# example 2
$s2 = join(',', $a);
# print
var_dump($s == 'May,June', $s2 == 'May,June');
