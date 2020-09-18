<?php
$s = 'C:\\Windows\\write.exe';
# example 1
$s1 = pathinfo($s, PATHINFO_DIRNAME);
# example 2
$o = new SplFileInfo($s);
$s2 = $o->getPath();
# print
var_dump($s1 == 'C:\\Windows', $s2 == 'C:\\Windows');
