<?php
$dir_o = new RecursiveDirectoryIterator('.');
$dir_o->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$iter_o = new RecursiveIteratorIterator($dir_o);

foreach ($iter_o as $info_o) {
   echo $info_o->getPathname(), "\n";
}
