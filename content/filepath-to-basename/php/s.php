<?php
$s = 'C:\\Windows\\notepad.exe';
# example 1
$s1 = basename($s);
# example 2
$s2 = pathinfo($s, PATHINFO_BASENAME);
# example 3
$s3 = basename($s, '.exe');
# example 4
$s4 = pathinfo($s, PATHINFO_FILENAME);
# print
var_dump(
   $s1 == 'notepad.exe',
   $s2 == 'notepad.exe',
   $s3 == 'notepad',
   $s4 == 'notepad'
);
