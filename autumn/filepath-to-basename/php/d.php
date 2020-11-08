<?php
$o = new SplFileInfo('C:\Windows\notepad.exe');
$s = $o->getFilename();
var_dump($s == 'notepad.exe');
