<?php
$a = ['May', 'June'];
# example 1
$s1 = implode(',', $a);
# example 2
$s2 = join(',', $a);
# print
var_dump($s1 == 'May,June', $s2 == 'May,June');
