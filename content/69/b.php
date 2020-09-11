<?php
$o = new SplFileInfo('index.md');
# example 1
$b = $o->isFile();
# example 2
$b2 = ! $o->isDir();
# print
var_dump($b, $b2);
