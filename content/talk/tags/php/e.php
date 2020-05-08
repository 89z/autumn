<?php
$s1 = realpath('index.md');
symlink($s1, '/cygdrive/c/index.md');
