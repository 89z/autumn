<?php
# example 1
$s1 = file_get_contents('a.tar.gz');
$s2 = gzdecode($s1);
file_put_contents('a.tar', $s2);
# example 2
copy('compress.zlib://a.tar.gz', 'a.tar');
