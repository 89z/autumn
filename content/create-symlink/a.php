<?php
$s_old = getcwd();
$s_new = getenv('TMP');
# Permission denied
symlink($s_old . '/index.md', $s_new . '/index.md');
