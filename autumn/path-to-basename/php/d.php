<?php
$f = new SplFileInfo('C:\Windows\notepad.exe');
$s = $f->getFilename();
var_dump($s == 'notepad.exe');
