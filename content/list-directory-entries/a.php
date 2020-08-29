<?php
$o_iter = new FilesystemIterator('.');
foreach ($o_iter as $o_info) {
   echo $o_info->getPathname(), "\n";
}
