<?php
$s = 'C:\Windows\notepad.exe';
# example 1
$s1 = pathinfo($s, PATHINFO_BASENAME);
# example 2
$s2 = pathinfo($s, PATHINFO_FILENAME);
# print
var_dump($s1 == 'notepad.exe', $s2 == 'notepad');
