<?php
# example 1
unlink('a.txt');
# example 2
$s_tmp = getenv('TMP');
$o_dir = new RecursiveDirectoryIterator($s_tmp);
$o_iter = new RecursiveIteratorIterator($o_dir);
foreach ($o_iter as $o_path) {
   if (! $o_path->isDir()) {
      unlink($o_path->getPathname());
   }
}
