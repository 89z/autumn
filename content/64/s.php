<?php
$s = 'C:\\php\\license.txt';
# example 1
$s2 = basename($s);
# example 2
$s3 = pathinfo($s, PATHINFO_BASENAME);
# example 3
$s4 = basename($s, '.txt');
# example 4
$s5 = pathinfo($s, PATHINFO_FILENAME);
# print
var_dump(
   $s2 == 'license.txt',
   $s3 == 'license.txt',
   $s4 == 'license',
   $s5 == 'license'
);
