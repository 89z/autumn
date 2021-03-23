<?php

function directoryIter(string $path): object {
   function filterIter($info) {
      return $info->getFilename() == '.git' ? false : true;
   }
   $dir = new RecursiveDirectoryIterator($path);
   $dir->setFlags(RecursiveDirectoryIterator::SKIP_DOTS);
   $filter = new RecursiveCallbackFilterIterator($dir, 'filterIter');
   return new RecursiveIteratorIterator($filter);
}

$dir = directoryIter('..');

foreach ($dir as $info) {
   echo $info->getPathname(), "\n";
}
