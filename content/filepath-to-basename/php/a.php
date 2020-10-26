<?php
$s = 'C:\Windows\notepad.exe';
# example 1
$s1 = basename($s);
# example 2
$s2 = basename($s, '.exe');
# print
var_dump($s1 == 'notepad.exe', $s2 == 'notepad');
