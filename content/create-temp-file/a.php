<?php
# create and remove
$r1 = tmpfile();
$s1 = stream_get_meta_data($r1)['uri'];
# create only
$s2 = tempnam('', '');
$r2 = fopen($s2, 'w');
# print
var_dump($s1, $s2);
