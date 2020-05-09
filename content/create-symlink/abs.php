<?php
$s_pub = getenv('PUBLIC');
$s_sep = DIRECTORY_SEPARATOR;
# round 1
$s_old = realpath('abs.php');
symlink($s_old, $s_pub . $s_sep . 'Music' . $s_sep . 'abs.php');
# round 2
$s_old = realpath('chdir.php');
symlink($s_old, $s_pub . $s_sep . 'Videos' . $s_sep . 'chdir.php');
