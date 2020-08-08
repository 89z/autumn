<?php
declare(strict_types = 1);
# example 1
$a1 = glob('a.*');
# example 2
$o_file = new FilesystemIterator('.');
foreach ($o_file as $o_name) {
   $a2[] = $o_name->getFilename();
}
# example 3
$n = RecursiveDirectoryIterator::SKIP_DOTS;
$o_dir = new RecursiveDirectoryIterator('.', $n);
$o_iter = new RecursiveIteratorIterator($o_dir);
foreach ($o_iter as $o_name) {
   $a3[] = $o_name->getPathname();
}
# print
var_dump($a1, $a2, $a3);
