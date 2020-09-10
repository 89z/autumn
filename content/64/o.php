<?php
$o = new SplFileInfo('C:\\php\\license.txt');
# example 1
$s1 = $o->getBasename();
# example 2
$s2 = $o->getFilename();
# example 3
$s3 = $o->getBasename('.txt');
# print
var_dump($s1 == 'license.txt', $s2 == 'license.txt', $s3 == 'license');
