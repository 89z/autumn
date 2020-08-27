<?php
# example 1
$a1 = glob('*');
# example 2
$o_file = new FilesystemIterator('.');
foreach ($o_file as $o_name) {
   $a2[] = $o_name->getFilename();
}
# example 3
$o_dir = new RecursiveDirectoryIterator('.', FilesystemIterator::SKIP_DOTS);
$o_iter = new RecursiveIteratorIterator($o_dir);
foreach ($o_iter as $o_name) {
   $a3[] = $o_name->getFilename();
}
# example 4
$f_filter = fn ($o_file) => $o_file->getFilename() == '.git' ? false : true;
$o_dir = new RecursiveDirectoryIterator('.');
$o_filter = new RecursiveCallbackFilterIterator($o_dir, $f_filter);
$o_iter = new RecursiveIteratorIterator($o_filter);
foreach ($o_iter as $o_name) {
   $a4[] = $o_name->getFilename();
}
# print
var_dump($a1, $a2, $a3, $a4);
