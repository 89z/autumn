<?php
$f = new SplFileInfo('a.php');
$b = $f->isFile();
var_dump($b);
