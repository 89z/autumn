<?php
$o = new SplFileInfo('C:\\php\\.\\license.txt');
$s1 = $o->getRealPath();
$s2 = $o->getBasename();
$s3 = $o->getFilename();
$s4 = $o->getBasename('.txt');
$s5 = $o->getExtension();
