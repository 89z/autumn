<?php

function directoryIter(string $path): object {
   $dir = new RecursiveDirectoryIterator($path);
   $dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
   return new RecursiveIteratorIterator($dir);
}

$dir = directoryIter('..');

foreach ($dir as $info) {
   echo $info->getPathname(), "\n";
}
