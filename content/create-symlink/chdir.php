<?php
function f_join(...$a_elem) {
   return implode(DIRECTORY_SEPARATOR, $a_elem);
}
$s_wd = getcwd();
$s_pub = getenv('PUBLIC');
# round 1
$s_old = f_join($s_wd, 'abs.php');
$s_new = f_join($s_pub, 'Music');
chdir($s_new);
symlink($s_old, 'abs.php');
# round 2
$s_old = f_join($s_wd, 'chdir.php');
$s_new = f_join($s_pub, 'Videos');
chdir($s_new);
symlink($s_old, 'chdir.php');
