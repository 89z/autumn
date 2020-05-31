<?php
$o_dir = new RecursiveDirectoryIterator('.');
$o_iter = new RecursiveIteratorIterator($o_dir);
foreach ($o_iter as $o_path) {
   echo $o_path->getPathname(), "\n";
}
