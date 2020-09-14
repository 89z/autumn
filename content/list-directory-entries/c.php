<?php
$o_dir = new RecursiveDirectoryIterator('.');
$o_dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$f_filter = fn ($o_info) => $o_info->getFilename() == '.git' ? false : true;
$o_filter = new RecursiveCallbackFilterIterator($o_dir, $f_filter);
$o_iter = new RecursiveIteratorIterator($o_filter);
foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
