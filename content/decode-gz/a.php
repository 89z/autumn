<?php
extension_loaded('zlib') or die('php-zlib');
$s1 = file_get_contents('a.tar.gz');
$s2 = gzdecode($s1);
file_put_contents('a.tar', $s2);
