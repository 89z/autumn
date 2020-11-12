<?php
$dir_o = new RecursiveDirectoryIterator('.');
$dir_o->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$filter_f = fn ($info_o) => $info_o->getFilename() == '.git' ? false : true;
$filter_o = new RecursiveCallbackFilterIterator($dir_o, $filter_f);
$iter_o = new RecursiveIteratorIterator($filter_o);

foreach ($iter_o as $info_o) {
   echo $info_o->getPathname(), "\n";
}
