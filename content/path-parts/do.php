<?php
$o = new SplFileInfo('C:\\php\\');
$s = $o->getPathname();
var_dump($s == 'C:\\php');
