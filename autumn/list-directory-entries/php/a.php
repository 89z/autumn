<?php
$iter = new FilesystemIterator('.');

foreach ($iter as $info) {
   echo $info->getPathname(), "\n";
}
