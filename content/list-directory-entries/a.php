<?php
echo "example 1\n";
$o_iter = new FilesystemIterator('.');
foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
echo "\nexample 2\n";
$o_dir = new RecursiveDirectoryIterator('.');
$o_dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$o_iter = new RecursiveIteratorIterator($o_dir);
foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
echo "\nexample 3\n";
$o_dir = new RecursiveDirectoryIterator('.');
$o_dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
$f_filter = fn ($o_info) => $o_info->getFilename() == '.git' ? false : true;
$o_filter = new RecursiveCallbackFilterIterator($o_dir, $f_filter);
$o_iter = new RecursiveIteratorIterator($o_filter);
foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
