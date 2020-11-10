<?php
$o = new SplFileInfo('C:\Users');
$b = $o->isDir();
var_dump($b);
