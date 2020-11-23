<?php
$s = 'C:\Windows\notepad.exe';
# example 1
$s1 = pathinfo($s, PATHINFO_DIRNAME);
# example 2
$m = pathinfo($s);
$s2 = $m['dirname'];
# print
var_dump($s1 == 'C:\Windows', $s2 == 'C:\Windows');
