<?php
$o = new SplFileInfo('index.md');
$b = ! $o->isDir();
var_dump($b);
