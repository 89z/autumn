<?php
$n1 = time();
touch('index.md', $n1);
$n2 = filemtime('index.md');
var_dump($n1, $n2 == $n1);
