<?php
$s = 'C:\Windows\notepad.exe';

$f = new SplFileInfo($s);
$t = $f->getPath();

var_dump($t == 'C:\Windows');
