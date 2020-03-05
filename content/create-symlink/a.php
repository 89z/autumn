<?php
# example 1
$s1 = realpath('index.md');
symlink($s1, '/Desktop/index.md');
# exmaple 2
$s2 = getcwd();
symlink("$s2/index.md", '/Desktop/index.md');
