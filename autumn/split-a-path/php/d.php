<?php
$f = new SplFileInfo('a.php');
$s = $f->getExtension();
var_dump($s == 'php');
