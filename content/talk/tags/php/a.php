<?php
$s1 = realpath('index.md');
$s2 = getenv('HOMEDRIVE');
symlink($s1, $s2 . DIRECTORY_SEPARATOR . 'a.md');
