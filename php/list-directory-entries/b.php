<?php
$o_dir = new RecursiveDirectoryIterator('.');
$o_dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$o_iter = new RecursiveIteratorIterator($o_dir);

foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
