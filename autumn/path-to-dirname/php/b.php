<?php
$s = 'C:\Windows\notepad.exe';
$o = new SplFileInfo($s);
$s1 = $o->getPath();
var_dump($s1 == 'C:\Windows');
