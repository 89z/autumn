<?php
$iter_o = new FilesystemIterator('.');

foreach ($iter_o as $info_o) {
   echo $info_o->getPathname(), "\n";
}
