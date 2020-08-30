<?php
$s = 'C:\\php\\license.txt';
# example 1
$s1 = basename($s);
# example 2
$s2 = pathinfo($s, PATHINFO_BASENAME);
# example 3
$s3 = basename($s, '.txt');
# example 4
$s4 = pathinfo($s, PATHINFO_FILENAME);
# print
var_dump(
   $s1 == 'license.txt',
   $s2 == 'license.txt',
   $s3 == 'license',
   $s4 == 'license'
);
