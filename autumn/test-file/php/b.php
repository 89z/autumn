<?php
$o = new SplFileInfo('a.php');
$b = $o->isFile();
var_dump($b);
