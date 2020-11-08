<?php
$o = new SplFileInfo('index.md');
$b = $o->isFile();
var_dump($b);
