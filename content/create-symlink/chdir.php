<?php
$s_wd = getcwd();
$s_pub = getenv('PUBLIC');
$s_sep = DIRECTORY_SEPARATOR;
# round 1
chdir($s_pub . $s_sep . 'Music');
symlink($s_wd . $s_sep . 'abs.php', 'abs.php');
# round 2
chdir($s_pub . $s_sep . 'Videos');
symlink($s_wd . $s_sep . 'chdir.php', 'chdir.php');
