<?php
$s1 = realpath('index.md');
$s2 = getenv('HOMDRIVE');
chdir($s2);
symlink($s1, 'index.md');
