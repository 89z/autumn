<?php
$o = new SplFileInfo('a.php');
$s = $o->getExtension();
var_dump($s == 'php');
