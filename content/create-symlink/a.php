<?php
# example 1
$s1 = getcwd();
symlink($s1 . DIRECTORY_SEPARATOR . 'index.md', 'a.md');
# example 2
$s2 = realpath('index.md');
symlink($s2, 'b.md');
