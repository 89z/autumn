<?php
$o = new SplFileInfo('C:\\php\\.\\license.txt');
# example 1
$s1 = $o->getRealPath();
# example 2
$s2 = $o->getFilename();
# example 3
$s3 = $o->getBasename();
# example 4
$s4 = $o->getBasename('.txt');
# example 5
$s5 = $o->getExtension();
# print
var_dump(
   $s1 == 'C:\\php\\license.txt',
   $s2 == 'license.txt',
   $s3 == 'license.txt',
   $s4 == 'license',
   $s5 == 'txt'
);
