<?php
$n = time();
touch('index.md', $n);
$n2 = filemtime('index.md');
var_dump($n, $n2 == $n);
