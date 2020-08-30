<?php
$s = 'C:\\php\\.\\license.txt';
# example 1
$s1 = realpath($s);
# example 2
$s2 = basename($s);
# example 3
$s3 = pathinfo($s, PATHINFO_BASENAME);
# example 4
$s4 = basename($s, '.txt');
# example 5
$s5 = pathinfo($s, PATHINFO_FILENAME);
# example 6
$s6 = pathinfo($s, PATHINFO_EXTENSION);
# print
var_dump(
   $s1 == 'C:\\php\\license.txt',
   $s2 == 'license.txt',
   $s3 == 'license.txt',
   $s4 == 'license',
   $s5 == 'license',
   $s6 == 'txt'
);
